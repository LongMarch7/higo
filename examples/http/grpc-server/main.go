package main
import (
    "context"
    "flag"
    "github.com/LongMarch7/higo/app"
    "github.com/LongMarch7/higo/middleware"
    "github.com/LongMarch7/higo/middleware/prometheus"
    "github.com/LongMarch7/higo/middleware/zipkin"
    "github.com/LongMarch7/higo/service/web"
    "github.com/LongMarch7/higo/util/log"
    "google.golang.org/grpc/grpclog"
    _ "github.com/LongMarch7/higo-web/controller/admin"
)

func main() {
    etcdServer := flag.String("e","http://localhost:8500","etcd service addr")
    prefix := flag.String("n","WebServer","prefix value")
    serviceAddress := flag.String("s","127.0.0.1","server addr")
    servicePort := flag.Int("p",0,"server port")
    threadMax := flag.String("c","1024","server thread pool max thread count")
    flag.Parse()
    ctx := context.Background()

    grpclog.SetLoggerV2(zap.NewDefaultLoggerConfig().NewLogger())


    server := app.NewServer(
        app.SConsulAddr(*etcdServer),
        app.SPrefix(*prefix),
        app.SServerAddr(*serviceAddress),
        app.SServerPort(*servicePort),
        app.SCtx(ctx),
        app.SMaxThreadCount(*threadMax),
        app.SzOptions([]zipkin.ZOption{ zipkin.Name(*prefix)}),
    )

    webServer := &web.GrpcServer{}
    webService := web.NewService()
    manager := middleware.NewMiddleware()
    webServer.HtmlCallHandler = manager.AddMiddleware(
        middleware.Prefix(*prefix),
        middleware.MethodName("HTML"),
        middleware.Endpoint(web.MakeHtmlCallServerEndpoint(webService)),
        middleware.POptions([]prometheus.POption{ prometheus.Name("HTML")}),
    ).NewServer()
    webServer.ApiCallHandler = manager.AddMiddleware(
        middleware.Prefix(*prefix),
        middleware.MethodName("API"),
        middleware.Endpoint(web.MakeApiCallServerEndpoint(webService)),
        middleware.POptions([]prometheus.POption{ prometheus.Name("API")}),
    ).NewServer()

    server.RegisterServiceServer(web.MakeRegisteFunc(webServer))
    server.Run()
}
