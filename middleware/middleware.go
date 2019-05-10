package middleware

import (
	"github.com/LongMarch7/higo/middleware/hystrix"
	"github.com/LongMarch7/higo/middleware/logger"
	"github.com/LongMarch7/higo/middleware/prometheus"
	"github.com/LongMarch7/higo/middleware/ratelimit"
	"github.com/LongMarch7/higo/middleware/zipkin"
	"github.com/LongMarch7/higo/tansport"
	"github.com/go-kit/kit/endpoint"
	"context"
	"errors"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"strings"
)
type Middleware struct {
	opts       middlewareOpt
	prometheus *prometheus.Prometheus
	endpoint   endpoint.Endpoint
}

func defaultConfig() middlewareOpt{
	return middlewareOpt{
		endpoint: nil,
		methodName: "default",
		prefix: "default",
		encodeFun: tansport.DefaultGrpcEncodeResponse,
		decodeFun: tansport.DefaultGrpcDecodeRequest,
	}
}

func NewMiddleware(opts ...MOption) *Middleware{
	opt := defaultConfig()
	for _, o := range opts {
		o(&opt)
	}
	return &Middleware{
		opts: opt,
	}
}

func (m *Middleware)AddMiddleware(opts ...MOption) *Middleware{
	for _, o := range opts {
		o(&m.opts)
	}
	var endpoint endpoint.Endpoint
	if m.opts.endpoint != nil {
		endpoint = m.opts.endpoint
		endpoint = ratelimit.NewLimiter(m.opts.rOptions...).Middleware()(endpoint)

		hOptions := append([]hystrix.HOption{},hystrix.Name(m.opts.methodName))
		hOptions = append(hOptions, m.opts.hOptions...)
		endpoint = hystrix.NewHystrix(hOptions...).Middleware()(endpoint)

		zOptions := append([]zipkin.ZOption{},zipkin.MethodName(m.opts.methodName))
		zOptions = append(zOptions, zipkin.Name(m.opts.prefix))
		zOptions = append(zOptions, m.opts.zOptions...)
		endpoint = zipkin.NewZipkin().Middleware(zOptions...)(endpoint)

		lOptions := append([]logger.LOption{}, logger.MethodName(m.opts.methodName))
		lOptions = append(lOptions, logger.Prefix(m.opts.prefix))
		lOptions = append(lOptions, m.opts.lOptions...)
		endpoint = logger.NewLogger().Middleware(lOptions...)(endpoint)

		pOptions := append([]prometheus.POption{}, prometheus.Subsystem(m.opts.prefix))
		pOptions = append(pOptions, prometheus.Name(m.opts.methodName))
		pOptions = append(pOptions, m.opts.pOptions...)
		m.prometheus = prometheus.NewPrometheus(pOptions...)
		lvs := []string{"method", m.opts.methodName,"error"}
		endpoint = m.prometheus.Middleware(prometheus.Lvs(lvs),
			prometheus.Class(prometheus.Counter_TYPE),
			prometheus.Name(strings.Split(prometheus.GetName(m.prometheus),"_")[0]+"_count"),
			prometheus.Help("Number of requests received"),
		)(endpoint)
		endpoint = m.prometheus.Middleware(prometheus.Lvs(lvs),
			prometheus.Class(prometheus.Histogram_TYPE),
			prometheus.Name(strings.Split(prometheus.GetName(m.prometheus),"_")[0]+"_latency_seconds"),
			prometheus.Help("Total duration of requests in seconds."),
		)(endpoint)
		m.endpoint = endpoint
	}
	return m
}

func (m *Middleware)Endpoint() endpoint.Endpoint {
	return m.endpoint
}

func (m *Middleware)NewServer() *grpc_transport.Server {
	if m.endpoint != nil {
		return grpc_transport.NewServer(
			m.endpoint,
			m.opts.decodeFun,
			m.opts.encodeFun,
		)
	}else {
		return grpc_transport.NewServer(
			func(ctx context.Context, request interface{}) (response interface{}, err error){
				return nil, errors.New("[m] not found")
			},
			m.opts.decodeFun,
			m.opts.encodeFun,
		)
	}
}