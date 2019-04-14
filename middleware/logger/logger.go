package logger


import (
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
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				grpclog.Info(h.opts.methodName, "--", err, "--", time.Since(begin))
			}(time.Now())
			return next(ctx, request)

		}
	}
}
