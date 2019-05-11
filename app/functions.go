package app

import (
    "github.com/opentracing/opentracing-go"
    "github.com/opentracing/opentracing-go/log"
)

func SpanDecoratorFunc(span opentracing.Span, method string, req, resp interface{}, grpcError error){
    span.LogFields(log.Object("gRPC request", req))
}
