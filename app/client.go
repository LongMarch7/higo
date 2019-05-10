package app

import (
    "context"
    "errors"
    zaplog "github.com/LongMarch7/go-web/plugin/zap-log"
    "github.com/LongMarch7/higo/middleware"
    "github.com/LongMarch7/higo/middleware/zipkin"
    "github.com/LongMarch7/higo/service/base"
    base_context "github.com/LongMarch7/higo/base"
    local_transport "github.com/LongMarch7/higo/tansport"
    "github.com/LongMarch7/higo/tansport/pool"
    "github.com/LongMarch7/higo/util/sd/consul"
    "github.com/go-kit/kit/endpoint"
    "github.com/go-kit/kit/sd"
    "github.com/go-kit/kit/sd/lb"
    grpc_transport "github.com/go-kit/kit/transport/grpc"
    "github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
    "github.com/hashicorp/consul/api"
    "google.golang.org/grpc"
    "google.golang.org/grpc/metadata"
    "github.com/grpc-ecosystem/go-grpc-middleware"
    "io"
    "time"
)

type serviceList struct{
    endpoint endpoint.Endpoint
    defaultEndpointer *sd.DefaultEndpointer
}
type Client struct {
    opts                  ClientOpt
    serviceList           map[string] serviceList
    zipkin                *zipkin.Zipkin
    dialOpts              []grpc.DialOption
}

func defaultClientConfig() ClientOpt{
    return ClientOpt{
        consulAddr: "http://localhost:8500",
        serviceName: "defaultServer",
        retryTime: time.Second * 3,
        retryCount: 3,
        factory: nil,
        passingOnly: true,
        logger: zaplog.NewDefaultLogger(),
        middleware: nil,
        encodeFunc: local_transport.DefaultGrpcEncodeRequestFunc,
        decodeFunc: local_transport.DefaultGrpcDecodeResponseFunc,
    }
}

func NewClient(opts ...COption) *Client{
    opt := defaultClientConfig()
    for _, o := range opts {
        o(&opt)
    }
    client := &Client{
        opts: opt,
        serviceList: make(map[string]serviceList),
    }
    client.init()
    pool.Init()
    return client
}

func (c *Client)init(){

}

func (c *Client)makeDefaultFactory() sd.Factory{
    return func(instance string) (endpoint.Endpoint, io.Closer, error) {
        return func(ctx context.Context, request interface{}) (interface{}, error) {
            poolManage,ok := pool.GetConnect(instance)
            if ! ok {
                return nil,errors.New("[p] not found")
            }

            cManager, err := pool.GetConnectFromPool(instance, poolManage, c.dialOpts...)
            if err != nil {
                return nil,err
            }
            defer func() {
                pool.PutConnectToPool(cManager, poolManage)
            }()
            par := ctx.Value("parameter")
            if par == nil {
                return nil,errors.New("parametes error")
            }
            parameter := par.(base.GrpcClientParameter)
            grpcEndpoint := grpc_transport.NewClient(
                cManager.Conn,
                parameter.Srv,
                parameter.Method,
                c.opts.encodeFunc,
                c.opts.decodeFunc,
                parameter.NewRlyFunc(),
                grpc_transport.ClientAfter(GrpcClientAfter),
            ).Endpoint()

            return grpcEndpoint(ctx,request)
        },nil,nil
    }
}

func (c *Client)AddEndpoint(opts ...COption){
    for _, o := range opts {
        o(&c.opts)
    }
    var client consul.Client
    {
        consulConfig := api.DefaultConfig()

        consulConfig.Address = c.opts.consulAddr
        consulClient, err := api.NewClient(consulConfig)
        if err != nil {
            c.opts.logger.Log("api.NewClient error")
            return
        }
        client = consul.NewClient(consulClient)
    }

    //创建实例管理器, 此管理器会Watch监听etc中prefix的目录变化更新缓存的服务实例数据
    consulTag := []string{"MicroServer",c.opts.serviceName}
    instancer := consul.NewInstancer(client, c.opts.logger, c.opts.serviceName, consulTag, c.opts.passingOnly, pool.Update)//pool.Update

    //创建端点管理器， 此管理器根据Factory和监听的到实例创建endPoint并订阅instancer的变化动态更新Factory创建的endPoint
    if c.opts.factory == nil{
        c.opts.factory = c.makeDefaultFactory()
    }
    endpointer := sd.NewEndpointer(instancer, c.opts.factory,  c.opts.logger)

    //创建负载均衡器
    balancer := lb.NewRoundRobin(endpointer)

    reqEndPoint := lb.Retry(c.opts.retryCount, c.opts.retryTime, balancer)

    if( c.opts.middleware != nil ){
        reqEndPoint = c.opts.middleware.AddMiddleware(middleware.Endpoint(reqEndPoint)).Endpoint()
    }

    zip := zipkin.NewZipkin()
    c.zipkin = zip
    dialOpts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
    if tracer := zip.GetTracer(); tracer != nil {
        dialOpts = append(dialOpts,grpc.WithUnaryInterceptor(
            grpc_middleware.ChainUnaryClient(
                withReqData(),
                otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads()),
            ),
        ))
    }
    c.dialOpts = dialOpts

    c.serviceList[c.opts.serviceName] = serviceList{
        endpoint: reqEndPoint,
        defaultEndpointer: endpointer,
    }
}


func (c *Client)GetClientEndpoint(srv string) endpoint.Endpoint{
    if service, ok := c.serviceList[srv]; ok{
        return service.endpoint
    }
    return func(ctx context.Context, request interface{}) (response interface{}, err error){
        c.opts.logger.Log(srv,"nothing done")
        return nil, errors.New("[c] not found")
    }
}

func withReqData() grpc.UnaryClientInterceptor{
    return func(
        ctx context.Context,
        method string,
        req, resp interface{},
        cc *grpc.ClientConn,
        invoker grpc.UnaryInvoker,
        opts ...grpc.CallOption,
    ) error {
        baseCtx := ctx.Value(base_context.StrucName)
        if baseCtx == nil {
            return errors.New("get context error")
        }
        baseContext := baseCtx.(*base_context.BaseContext)
        md, ok := metadata.FromOutgoingContext(ctx)
        if !ok {
            md = metadata.New(nil)
        } else {
            md = md.Copy()
        }
        reqParams := baseContext.Params
        for key, value := range reqParams {
            md[key] = []string{value}
        }
        ctxWithMetadata := metadata.NewOutgoingContext(ctx, md)
        return invoker(ctxWithMetadata,method,req,resp,cc,opts...)
    }
}

func GrpcClientAfter(ctx context.Context, header metadata.MD, trailer metadata.MD) context.Context{
    baseCtx := ctx.Value(base_context.StrucName)
    if baseCtx != nil {
        baseContext := baseCtx.(*base_context.BaseContext)
        if len(header) > 0{
            baseContext.GrpcHeader = header
        }
        if len(trailer) > 0{
            baseContext.GrpcTrailer = trailer
        }
    }

    return ctx
}