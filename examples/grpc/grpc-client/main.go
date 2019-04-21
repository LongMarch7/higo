package main

import (
	"flag"
	"github.com/LongMarch7/higo/middleware"
	"github.com/LongMarch7/higo/service/admin/setting"
	"google.golang.org/grpc/grpclog"
	"github.com/LongMarch7/higo/app"
	"github.com/LongMarch7/higo/util/log"
	"os"
	"sync"
	"time"
	"context"
)

var c chan os.Signal
var wg sync.WaitGroup

func Producer(){
Loop:
	for{
		select {
		case s := <-c:
			grpclog.Info("Producer get", s)
			break Loop
		default:
		}
		time.Sleep(500 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	etcdServer := flag.String("e","http://localhost:8500","etcd service addr")
	prefix := flag.String("p","sayHelloService","prefix value")
	flag.Parse()


	grpclog.SetLoggerV2(zap.NewDefaultLoggerConfig().NewLogger())

	mw := middleware.NewMiddleware(middleware.Prefix("gateway"),middleware.MethodName("request"))
	client := app.NewClient(
		app.CConsulAddr(*etcdServer),
		app.CPrefix(*prefix),
		app.CRetryCount(3),
		app.CMiddleware(mw),
		app.CRetryTime(time.Second * 3),
	)
	setting.SayHelloProxy(client.GetClientEndpoint("sayHelloService"))(context.Background(),"jack")
}

