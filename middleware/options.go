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
type middlewareOpt struct {
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
type MOption func(o *middlewareOpt)

func ZOptions(zOptions []zipkin.ZOption) MOption {
	return func(o *middlewareOpt) {
		o.zOptions = zOptions
	}
}

func ROptions(rOptions []ratelimit.ROption) MOption {
	return func(o *middlewareOpt) {
		o.rOptions = rOptions
	}
}

func POptions(pOptions []prometheus.POption) MOption {
	return func(o *middlewareOpt) {
		o.pOptions = pOptions
	}
}

func HOptions(hOptions []hystrix.HOption) MOption {
	return func(o *middlewareOpt) {
		o.hOptions = hOptions
	}
}

func Endpoint(endpoint endpoint.Endpoint) MOption {
	return func(o *middlewareOpt) {
		o.endpoint = endpoint
	}
}

func Prefix(prefix string) MOption {
	return func(o *middlewareOpt) {
		o.prefix = prefix
	}
}

func MethodName(methodName string) MOption {
	return func(o *middlewareOpt) {
		o.methodName = methodName
	}
}

func DecodeFun(decodeFun grpc_transport.DecodeRequestFunc) MOption {
	return func(o *middlewareOpt) {
		o.decodeFun = decodeFun
	}
}

func EncodeFun(encodeFun grpc_transport.EncodeResponseFunc) MOption {
	return func(o *middlewareOpt) {
		o.encodeFun = encodeFun
	}
}

