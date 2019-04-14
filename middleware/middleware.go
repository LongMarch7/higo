package middleware

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/LongMarch7/higo/middleware/ratelimit"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"github.com/LongMarch7/higo/middleware/hystrix"
	"github.com/LongMarch7/higo/middleware/zipkin"
	"github.com/LongMarch7/higo/middleware/logger"
	"github.com/LongMarch7/higo/middleware/prometheus"
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

func NewMiddlewareManager() *MiddlewareManager{
	opt := defaultConfig()

	middleware := &MiddlewareManager{
		opts: opt,
	}
	return middleware
}

func (m *MiddlewareManager)AddService(opts ...MOption) *grpc_transport.Server{
	for _, o := range opts {
		o(&m.opts)
	}
	var pluginEndpoint endpoint.Endpoint
	if m.opts.endpoint != nil {
		pluginEndpoint = m.opts.endpoint
		pluginEndpoint = ratelimit.NewLimiter(m.opts.rOptions...).Middleware()(pluginEndpoint)
		pluginEndpoint = hystrix.NewHystrix(m.opts.hOptions...).Middleware()(pluginEndpoint)
		pluginEndpoint = zipkin.NewZipkin(m.opts.zOptions...).Middleware(zipkin.Name(m.opts.methodName))(pluginEndpoint)
		pluginEndpoint = logger.NewLogger(m.opts.lOptions...).Middleware()(pluginEndpoint)
		m.prometheus = prometheus.NewPrometheus(m.opts.pOptions...)
		pluginEndpoint = m.prometheus.Middleware()(pluginEndpoint)
		lvs := []string{"method", m.opts.methodName,"error"}
		m.prometheus.AddObj(prometheus.Lvs(lvs),prometheus.Class(prometheus.Counter_TYPE))
		m.prometheus.AddObj(prometheus.Lvs(lvs),prometheus.Class(prometheus.Histogram_TYPE))
		server := grpc_transport.NewServer(
			pluginEndpoint,
			m.opts.decodeFun,
			m.opts.encodeFun,
		)
		return server
	}else{
		return nil
	}
}