package zipkin

import (
	"encoding/json"
	"github.com/LongMarch7/higo/util/define"
	"github.com/go-kit/kit/endpoint"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc/grpclog"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"
	otext "github.com/opentracing/opentracing-go/ext"
	"context"
	"strings"
	"sync"
	"github.com/LongMarch7/higo/base"
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
	//return kitopentracing.TraceClient(z.zip, z.opts.methodName)
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			var clientSpan stdopentracing.Span
			method :=z.opts.methodName
			methodNameByCtx := ctx.Value(define.PatternName)
			if methodNameByCtx != nil {
				method = methodNameByCtx.(string)
			}
			if parentSpan := stdopentracing.SpanFromContext(ctx); parentSpan != nil {
				clientSpan = z.zip.StartSpan(
					method,
					stdopentracing.ChildOf(parentSpan.Context()),
				)
			} else {
				clientSpan = z.zip.StartSpan(method)
			}
			baseCtx := ctx.Value(define.StrucName)
			if baseCtx != nil {
				pStrings, pErr := json.Marshal(baseCtx.(*base.BaseContext).Params)
				if pErr == nil {
					clientSpan.LogFields(log.String("gRPC req", string(pStrings)))
				}
			}
			defer func() {
				if err == nil{
					value,ok := base.GetDataFromGrpcResHeader(ctx,define.ResTypeName)
					if ok{
						if strings.Compare(value,"html") != 0 {
							clientSpan.LogFields(log.Object("gRPC res", response))
						}
					}
				}
				clientSpan.Finish()
			}()
			otext.SpanKindRPCClient.Set(clientSpan)
			ctx = stdopentracing.ContextWithSpan(ctx, clientSpan)
			return next(ctx, request)
		}
	}
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
