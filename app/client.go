package app

import (
    "context"
    "errors"
    zaplog "github.com/LongMarch7/go-web/plugin/zap-log"
    "github.com/LongMarch7/higo/middleware"
    "github.com/LongMarch7/higo/middleware/zipkin"
    "github.com/LongMarch7/higo/tansport/pool"
    "github.com/LongMarch7/higo/util/sd/consul"
    "github.com/go-kit/kit/endpoint"
    "github.com/go-kit/kit/sd"
    "github.com/go-kit/kit/sd/lb"
    "github.com/go-kit/kit/transport/grpc/_grpc_test/pb"
    "github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
    "github.com/hashicorp/consul/api"
    "google.golang.org/grpc"
    grpc_transport "github.com/go-kit/kit/transport/grpc"
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
    grpcReply             interface{}
}

func defaultClientConfig() ClientOpt{
    return ClientOpt{
        consulAddr: "http://localhost:8500",
        prefix: "bookServer",
        retryTime: time.Second * 3,
        retryCount: 3,
        passingOnly: true,
        logger: zaplog.NewDefaultLogger(),
        middle: nil,
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
    zip := zipkin.NewZipkin(zipkin.Name(c.opts.prefix))
    c.zipkin = zip
    dialOpts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
    if tracer := zip.GetTracer(); tracer != nil {
        dialOpts = append(dialOpts,grpc.WithUnaryInterceptor(
            otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads()),
        ))
    }
    c.dialOpts = dialOpts
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
            grpc_transport.NewClient(
                cManager.Conn,
                ctx.Value("srv").(string),
                ctx.Value("method").(string),
                c.opts.encodeFunc,
                c.opts.decodeFUnc,
                pb.TestResponse{},
            ).Endpoint()
            return nil,err
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
    consulTag := []string{c.opts.prefix}
    instancer := consul.NewInstancer(client, c.opts.logger, c.opts.prefix, consulTag, c.opts.passingOnly, pool.Update)//pool.Update

    //创建端点管理器， 此管理器根据Factory和监听的到实例创建endPoint并订阅instancer的变化动态更新Factory创建的endPoint
    endpointer := sd.NewEndpointer(instancer, c.opts.factory,  c.opts.logger)

    //创建负载均衡器
    balancer := lb.NewRoundRobin(endpointer)

    reqEndPoint := lb.Retry(c.opts.retryCount, c.opts.retryTime, balancer)

    if( c.opts.middle != nil ){
        reqEndPoint = c.opts.middle.AddMiddleware(middleware.Endpoint(reqEndPoint)).Endpoint()
    }
    c.serviceList[c.opts.prefix] = serviceList{
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