package main

//import (
//	"context"
//	"flag"
//	"github.com/LongMarch7/higo/app"
//	"github.com/LongMarch7/higo/middleware"
//	"github.com/LongMarch7/higo/service/examples/test"
//	"github.com/LongMarch7/higo/util/log"
//	"google.golang.org/grpc/grpclog"
//	"os"
//	"sync"
//	"time"
//)
//
//var c chan os.Signal
//var wg sync.WaitGroup
//
//func Producer(){
//Loop:
//	for{
//		select {
//		case s := <-c:
//			grpclog.Info("Producer get", s)
//			break Loop
//		default:
//		}
//		time.Sleep(500 * time.Millisecond)
//	}
//	wg.Done()
//}
//
//func main() {
//	etcdServer := flag.String("e","http://localhost:8500","etcd service addr")
//	flag.Parse()
//
//
//	grpclog.SetLoggerV2(zap.NewDefaultLoggerConfig().NewLogger())
//
//	mw := middleware.NewMiddleware(middleware.Prefix("gateway"),middleware.MethodName("request"))
//	client := app.NewClient(
//		app.CConsulAddr(*etcdServer),
//		app.CRetryCount(3),
//		app.CRetryTime(time.Second * 3),
//	)
//	serviceName := "SettingServer"
//	client.AddEndpoint(app.CMiddleware(mw),app.CServiceName(serviceName))
//	for {
//		rs, err := test.SayHelloProxy(client.GetClientEndpoint(serviceName))(context.Background(),&test.TestStrucAlias{Test1: "jack"})
//		grpclog.Info("[rs]=",rs," [err]=",err," [len(err)]",len(err))
//		time.Sleep(time.Second)
//	}
//}

