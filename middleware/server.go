package middleware

import (
	"github.com/LongMarch7/higo/middleware/hystrix"
	"github.com/LongMarch7/higo/middleware/logger"
	"github.com/LongMarch7/higo/middleware/prometheus"
	"github.com/LongMarch7/higo/middleware/ratelimit"
	"github.com/LongMarch7/higo/middleware/zipkin"
	"github.com/LongMarch7/higo/tansport"
	"github.com/go-kit/kit/endpoint"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
)
type MiddlewareServer struct {
	opts       middlewareServerOpt
	prometheus *prometheus.Prometheus
}

func defaultServerConfig() middlewareServerOpt{
	return middlewareServerOpt{
		endpoint: nil,
		methodName: "default",
		prefix: "default",
		encodeFun: tansport.DefaultGrpcEncodeResponse,
		decodeFun: tansport.DefaultGrpcDecodeRequest,
	}
}

func NewServerMiddleware() *MiddlewareServer{
	opt := defaultServerConfig()
	return &MiddlewareServer{
		opts: opt,
	}
}

func (m *MiddlewareServer)AddMiddleware(opts ...SMOption) *MiddlewareServer{
	for _, o := range opts {
		o(&m.opts)
	}
	var endpoint endpoint.Endpoint
	if m.opts.endpoint != nil {
		endpoint = m.opts.endpoint
		endpoint = ratelimit.NewLimiter(m.opts.rOptions...).Middleware()(endpoint)

		endpoint = hystrix.NewHystrix(m.opts.hOptions...).Middleware()(endpoint)

		zOptions := append([]zipkin.ZOption{},zipkin.MethodName(m.opts.methodName))
		zOptions = append(zOptions, zipkin.Name(m.opts.prefix))
		zOptions = append(zOptions, m.opts.zOptions...)
		endpoint = zipkin.NewZipkin(zOptions...).Middleware(zipkin.Name(m.opts.methodName))(endpoint)

		lOptions := append([]logger.LOption{}, logger.MethodName(m.opts.methodName))
		lOptions = append(lOptions, m.opts.lOptions...)
		endpoint = logger.NewLogger(lOptions...).Middleware()(endpoint)

		pOptions := append([]prometheus.POption{}, prometheus.Subsystem(m.opts.prefix))
		pOptions = append([]prometheus.POption{}, prometheus.Name(m.opts.methodName))
		pOptions = append(pOptions, m.opts.pOptions...)
		m.prometheus = prometheus.NewPrometheus(pOptions...)
		lvs := []string{"method", m.opts.methodName,"error"}
		endpoint = m.prometheus.Middleware(prometheus.Lvs(lvs),
			prometheus.Class(prometheus.Counter_TYPE),
			prometheus.Name(prometheus.GetName(m.prometheus)+"_count"),
			prometheus.Help("Number of requests received"),
		)(endpoint)
		endpoint = m.prometheus.Middleware(prometheus.Lvs(lvs),
			prometheus.Class(prometheus.Histogram_TYPE),
			prometheus.Name(prometheus.GetName(m.prometheus)+"_latency_seconds"),
			prometheus.Help("Total duration of requests in seconds."),
		)(endpoint)
		m.opts.endpoint = endpoint
	}
	return m
}

func (m *MiddlewareServer)NewServer() *grpc_transport.Server {
	if m.opts.endpoint != nil {
		return grpc_transport.NewServer(
			m.opts.endpoint,
			m.opts.decodeFun,
			m.opts.encodeFun,
		)
	}else {
		return nil
	}
}