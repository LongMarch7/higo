package main

import (
    "context"
    "flag"
    "github.com/LongMarch7/higo/app"
    "github.com/LongMarch7/higo/middleware"
    "github.com/LongMarch7/higo/middleware/prometheus"
    "github.com/LongMarch7/higo/service/admin/setting"
    "github.com/LongMarch7/higo/util/log"
    "google.golang.org/grpc/grpclog"
)

func main() {
    etcdServer := flag.String("e","http://localhost:8500","etcd service addr")
    prefix := flag.String("n","SettingServer","prefix value")
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
    )
    settingServer := &setting.GrpcServer{}
    manager := middleware.NewMiddleware()
    settingServer.SayHelloHandler = manager.AddMiddleware(
        middleware.Prefix(*prefix),
        middleware.MethodName("SettingServer"),
        middleware.Endpoint(setting.MakeSayHelloServerEndpoint(setting.NewService())),
        middleware.POptions([]prometheus.POption{ prometheus.Name("SayHello")}),
        ).NewServer()
    settingServer.DeleteuserHandler = manager.AddMiddleware(
        middleware.Prefix(*prefix),
        middleware.MethodName("SettingServer"),
        middleware.Endpoint(setting.MakeDeleteuserServerEndpoint(setting.NewService())),
        middleware.POptions([]prometheus.POption{ prometheus.Name("Deleteuser")}),
    ).NewServer()
    server.RegisterServiceServer(setting.MakeRegisteFunc(settingServer))
    server.Run()
}