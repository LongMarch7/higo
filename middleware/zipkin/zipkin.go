package zipkin

import (
	"github.com/go-kit/kit/endpoint"
	kitopentracing "github.com/go-kit/kit/tracing/opentracing"
	"google.golang.org/grpc/grpclog"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"
	"context"
	"sync"
)

type Zipkin struct{
	opts zipkinOpt
	collector zipkinot.Collector
	zip stdopentracing.Tracer
	inited bool
}

var initOpt sync.Once
var zipkin *Zipkin

func defaultConfig() zipkinOpt{
	return zipkinOpt{
		name:      "default",
		url:       "http://127.0.0.1:9411/api/v1/spans",
		hostPort:  "localhost:0",
		debug:     false,
		methodName: "default",
	}
}


func NewZipkin(opts ...ZOption) *Zipkin{
	initOpt.Do(func() {
		opt := defaultConfig()
		for _, o := range opts {
			o(&opt)
		}
		zipkin = &Zipkin{
			opts: opt,
			inited: false,
		}
		zipkin.zip , zipkin.collector = zipkin.Init()
	})
	return zipkin
}

func (z *Zipkin)Middleware(opts ...ZOption) endpoint.Middleware {
	for _, o := range opts {
		o(&z.opts)
	}
	if z.zip == nil {
		grpclog.Error("NewZipkinTracerMiddleware ", z.opts.name)
		return  func(next endpoint.Endpoint) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (interface{}, error){
				return next(ctx, request)
			}
		}
	}
	return kitopentracing.TraceClient(z.zip, z.opts.methodName)
}


func (z *Zipkin)Init() (stdopentracing.Tracer, zipkinot.Collector){
	local_collector, err := zipkinot.NewHTTPCollector(z.opts.url)
	if err != nil {
		grpclog.Error("zipkinot.NewHTTPCollector faild")
		return nil,nil
	}
	recorder := zipkinot.NewRecorder(local_collector, z.opts.debug, z.opts.hostPort, z.opts.name)
	newTracer, err := zipkinot.NewTracer(recorder)
	if err != nil {
		grpclog.Error("zipkinot.NewTracer faild", z.opts.name)
		return nil, local_collector
	}
	return newTracer,local_collector
}

func (z *Zipkin)GetTracer() stdopentracing.Tracer{
	return z.zip
}

func (z *Zipkin)Close(){
	if z.collector != nil {
		z.collector.Close()
		z.collector = nil
	}
}
