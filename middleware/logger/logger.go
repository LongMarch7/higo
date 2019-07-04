package logger


import (
	"github.com/LongMarch7/higo/util/define"
	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc/grpclog"
	"sync"
	"time"
	"context"
)

type Logger struct{
	opts LoggerOpt
}

var initOpt sync.Once
var logger *Logger

func defaultConfig() LoggerOpt{
	return LoggerOpt{
		methodName:   "default",
	}
}

func NewLogger(opts ...LOption) *Logger{
	initOpt.Do(func() {
		opt := defaultConfig()
		for _, o := range opts {
			o(&opt)
		}
		logger = &Logger{
			opts: opt,
		}
	})

	return logger
}

func (h *Logger)Middleware(opts ...LOption) endpoint.Middleware{
	for _, o := range opts {
		o(&h.opts)
	}
	prefix :=  h.opts.prefix
	method := h.opts.methodName
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func() {
				methodNameByCtx := ctx.Value(define.ReqPatternName)
				pattern :=""
				if methodNameByCtx != nil {
					pattern = methodNameByCtx.(string)
				}
				if err !=nil {
					grpclog.Error("[" + prefix + "|" + method + "|" + pattern + "]", err.Error(), "--", time.Now().Format("2006/1/2 15:04:05"))
				}else{
					grpclog.Info("[" + prefix + "|" + method + "|" + pattern + "]", "success--", time.Now().Format("2006/1/2 15:04:05"))
				}
			}()
			return next(ctx, request)

		}
	}
}
