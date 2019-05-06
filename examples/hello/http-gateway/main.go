package main

import (
	"flag"
	"github.com/LongMarch7/higo/app"
	"github.com/LongMarch7/higo/middleware"
	"github.com/LongMarch7/higo/router"
	"github.com/LongMarch7/higo/service/examples/hello"
	"github.com/LongMarch7/higo/util/log"
	"google.golang.org/grpc/grpclog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
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
	flag.Parse()


	grpclog.SetLoggerV2(zap.NewDefaultLoggerConfig().NewLogger())

	mw := middleware.NewMiddleware(middleware.Prefix("gateway"),middleware.MethodName("request"))
	client := app.NewClient(
		app.CConsulAddr(*etcdServer),
		app.CRetryCount(3),
		app.CRetryTime(time.Second * 3),
	)
	serviceName := "SettingServer"
	client.AddEndpoint(app.CMiddleware(mw),app.CServiceName(serviceName))

	r := router.NewRouter()

	r.Cli = client
	r.Add([]router.Routs{
		{"post|get","/test/{serviceName}",hello.MakeHelloWorldHandler(client.GetClientEndpoint(serviceName))},
	})
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir("E://go_project/higo/src/github.com/LongMarch7/higo/static"))))
	c = make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	wg.Add(1)
	go http.ListenAndServe("localhost:8080", r)
	go Producer()
	wg.Wait()
}

