package middleware

import (
	"github.com/LongMarch7/higo/middleware/zipkin"
	"github.com/LongMarch7/higo/middleware/ratelimit"
	"github.com/LongMarch7/higo/middleware/prometheus"
	"github.com/LongMarch7/higo/middleware/hystrix"
	"github.com/LongMarch7/higo/middleware/logger"
	"github.com/go-kit/kit/endpoint"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
)
type middlewareServerOpt struct {
	zOptions   []zipkin.ZOption
	rOptions   []ratelimit.ROption
	pOptions   []prometheus.POption
	hOptions   []hystrix.HOption
	lOptions   []logger.LOption
	endpoint   endpoint.Endpoint
	prefix     string
	methodName string
	decodeFun  grpc_transport.DecodeRequestFunc
	encodeFun  grpc_transport.EncodeResponseFunc
}
type SMOption func(o *middlewareServerOpt)

func SZOptions(zOptions []zipkin.ZOption) SMOption {
	return func(o *middlewareServerOpt) {
		o.zOptions = zOptions
	}
}

func SROptions(rOptions []ratelimit.ROption) SMOption {
	return func(o *middlewareServerOpt) {
		o.rOptions = rOptions
	}
}

func SPOptions(pOptions []prometheus.POption) SMOption {
	return func(o *middlewareServerOpt) {
		o.pOptions = pOptions
	}
}

func SHOptions(hOptions []hystrix.HOption) SMOption {
	return func(o *middlewareServerOpt) {
		o.hOptions = hOptions
	}
}

func SEndpoint(endpoint endpoint.Endpoint) SMOption {
	return func(o *middlewareServerOpt) {
		o.endpoint = endpoint
	}
}

func SPrefix(prefix string) SMOption {
	return func(o *middlewareServerOpt) {
		o.prefix = prefix
	}
}

func SMethodName(methodName string) SMOption {
	return func(o *middlewareServerOpt) {
		o.methodName = methodName
	}
}

func SDecodeFun(decodeFun grpc_transport.DecodeRequestFunc) SMOption {
	return func(o *middlewareServerOpt) {
		o.decodeFun = decodeFun
	}
}

func SEncodeFun(encodeFun grpc_transport.EncodeResponseFunc) SMOption {
	return func(o *middlewareServerOpt) {
		o.encodeFun = encodeFun
	}
}


type middlewareClientOpt struct {
	zOptions   []zipkin.ZOption
	rOptions   []ratelimit.ROption
	pOptions   []prometheus.POption
	hOptions   []hystrix.HOption
	lOptions   []logger.LOption
	endpoint   endpoint.Endpoint
	prefix     string
	methodName string
	decodeFun  grpc_transport.DecodeRequestFunc
	encodeFun  grpc_transport.EncodeResponseFunc
}
type CMOption func(o *middlewareClientOpt)

