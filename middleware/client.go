package middleware

import (
    "github.com/LongMarch7/higo/middleware/prometheus"
    "github.com/LongMarch7/higo/tansport"
)

type MiddlewareClient struct {
    opts       middlewareClientOpt
    prometheus *prometheus.Prometheus
}

func defaultClientConfig() middlewareClientOpt{
    return middlewareClientOpt{
        endpoint: nil,
        methodName: "default",
        prefix: "default",
        encodeFun: tansport.DefaultGrpcEncodeResponse,
        decodeFun: tansport.DefaultGrpcDecodeRequest,
    }
}

func NewSClientMiddleware() *MiddlewareClient{
    opt := defaultClientConfig()
    return &MiddlewareClient{
        opts: opt,
    }
}

