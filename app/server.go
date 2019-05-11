package app

import (
    "context"
    "github.com/LongMarch7/higo/controller/base"
    "github.com/LongMarch7/higo/util/define"
    "github.com/LongMarch7/higo/util/log"
    "github.com/gorilla/mux"
    "github.com/grpc-ecosystem/go-grpc-middleware"
    "github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
    "net"
    "net/http"
    "os"
    "os/signal"
    "strconv"
    "github.com/go-kit/kit/log"
    "github.com/LongMarch7/higo/middleware/zipkin"
    grpc_transport "github.com/go-kit/kit/transport/grpc"
    "sync"
    "time"
    "github.com/LongMarch7/higo/tansport"
)

type Server struct {
    opts                  ServerOpt
    listenConnector       net.Listener
    zipkin                *zipkin.Zipkin
    server                *grpc.Server
    c                     chan os.Signal
    wg                    sync.WaitGroup
}

func defaultServerConfig() ServerOpt{
    return ServerOpt{
        consulAddr: "http://localhost:8500",
        prefix: "bookServer",
        serverAddr: "127.0.0.1",
        serverPort: 0,
        ctx: context.Background(),
        maxThreadCount: "1024",
        netType: "tcp",
        advertiseAddress: "192.168.1.80",
        advertisePort: "10086",
        logger: zap.NewDefaultLogger(),
    }
}

func NewServer(opts ...SOption) *Server{
    opt := defaultServerConfig()
    for _, o := range opts {
        o(&opt)
    }
    server := &Server{
        opts: opt,
    }
    server.init()
    return server
}

func (s *Server)init(){
    ls, _ := net.Listen(s.opts.netType, s.opts.serverAddr+":"+strconv.Itoa(s.opts.serverPort))
    s.listenConnector = ls

    zip := zipkin.NewZipkin(s.opts.zOptions...)
    s.zipkin = zip
    var opts []grpc.ServerOption
    if tracer := zip.GetTracer(); tracer != nil {
        opts = append(opts,grpc_middleware.WithUnaryServerChain(
            otgrpc.OpenTracingServerInterceptor(tracer),
            ReqMetadaResolve,
        ),)
    }else{
        opts =[]grpc.ServerOption{grpc.UnaryInterceptor(grpc_transport.Interceptor)}
    }
    s.server = grpc.NewServer(opts...)
}

func  (s *Server)RegisterServiceServer(register define.GrpcRegister){
    register(s.server)
}

func (s *Server)Run(){
    s.c = make(chan os.Signal, 1)
    signal.Notify(s.c, os.Interrupt, os.Kill)
    s.wg.Add(1)

    port := s.listenConnector.Addr().(*net.TCPAddr).Port
    // 创建注册器
    config := RegisterConfig{
        consulAddress: s.opts.consulAddr,
        prefix: s.opts.prefix,
        service: s.opts.serverAddr,
        port: port,
        advertiseAddress: s.opts.advertiseAddress,
        advertisePort: s.opts.advertisePort,
        logger: log.NewNopLogger(),
        maxThreadCount: s.opts.maxThreadCount,
    }
    registrar := Register(config)
    registrar.Register()
    defer func(){
        s.zipkin.Close()
        registrar.Deregister()
        grpclog.Info("exit....")
        registrar = nil
        s.server.Stop()
    }()
    go func() {
        http.ListenAndServe(":" + s.opts.advertisePort, tansport.MakeMonitoringHttpHandler( mux.NewRouter(),tansport.MakeHealthEndpoint()))
    }()
    go s.server.Serve(s.listenConnector)
    go s.Producer()
    s.wg.Wait()
}

func (s *Server)Producer(){
Loop:
    for{
        select {
        case s := <-s.c:
            grpclog.Error("Producer | get", s)
            break Loop
        default:
        }
        time.Sleep(500 * time.Millisecond)
    }
    s.wg.Done()
}

func ReqMetadaResolve(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
    ctx = base.NewParameter(ctx)
    return handler(ctx, req)
}