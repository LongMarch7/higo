package middleware

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/LongMarch7/higo/middleware/ratelimit"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"github.com/LongMarch7/higo/middleware/hystrix"
	"github.com/LongMarch7/higo/middleware/zipkin"
	"github.com/LongMarch7/higo/middleware/logger"
	"github.com/LongMarch7/higo/middleware/prometheus"
	"sync"
)
type MiddlewareManager struct {
	opts       middlewareOpt
	prometheus *prometheus.Prometheus
}

func defaultConfig() middlewareOpt{
	return middlewareOpt{
		endpoint: nil,
		methodName: "default",
	}
}
var initOpt sync.Once
var middlewareManager *MiddlewareManager

func NewMiddlewareManager() *MiddlewareManager{
	initOpt.Do(func() {
		middlewareManager = &MiddlewareManager{}
	})
	opt := defaultConfig()
	middlewareManager.opts = opt
	return middlewareManager
}

func (m *MiddlewareManager)AddMiddleware(opts ...MOption){
	for _, o := range opts {
		o(&m.opts)
	}
	var endpoint endpoint.Endpoint
	if m.opts.endpoint != nil {
		endpoint = m.opts.endpoint
		endpoint = ratelimit.NewLimiter(m.opts.rOptions...).Middleware()(endpoint)

		endpoint = hystrix.NewHystrix(m.opts.hOptions...).Middleware()(endpoint)

		zOptions := append([]zipkin.ZOption{},zipkin.MethodName(m.opts.methodName))
		zOptions = append(zOptions, m.opts.zOptions...)
		endpoint = zipkin.NewZipkin(zOptions...).Middleware(zipkin.Name(m.opts.methodName))(endpoint)

		lOptions := append([]logger.LOption{}, logger.MethodName(m.opts.methodName))
		lOptions = append(lOptions, m.opts.lOptions...)
		endpoint = logger.NewLogger(lOptions...).Middleware()(endpoint)

		m.prometheus = prometheus.NewPrometheus(m.opts.pOptions...)
		lvs := []string{"method", m.opts.methodName,"error"}
		endpoint = m.prometheus.Middleware(prometheus.Lvs(lvs),prometheus.Class(prometheus.Counter_TYPE))(endpoint)
		endpoint = m.prometheus.Middleware(prometheus.Lvs(lvs),prometheus.Class(prometheus.Histogram_TYPE))(endpoint)
		m.opts.endpoint = endpoint
	}
}

func (m *MiddlewareManager)NewServer(opts ...MOption) *grpc_transport.Server {
	for _, o := range opts {
		o(&m.opts)
	}
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