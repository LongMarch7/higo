package ratelimit

import (
	"github.com/LongMarch7/higo/util/define"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
	"google.golang.org/grpc/grpclog"
	"sync"
	"time"
	"context"
)

type Limit struct{
	opts         ratelimitOpt
	rateLimitMap map[string]ratelimit.Allower
}

var initOpt sync.Once
var limit *Limit
func defaultConfig() ratelimitOpt{
	return ratelimitOpt{
		interval: time.Millisecond * 10,
		burst: 100,
	}
}

func NewLimiter(opts ...ROption) *Limit{
	initOpt.Do(func() {
		opt := defaultConfig()
		for _, o := range opts {
			o(&opt)
		}
		limit = &Limit{
			opts: opt,
			rateLimitMap: make(map[string]ratelimit.Allower),
		}
	})
	return limit
}

func (l *Limit)Middleware() endpoint.Middleware {
	return l.NewErroringLimiter()
}

func (l *Limit)AddRateLimitForMethod(name string){
	_,ok := l.rateLimitMap[name]
	if !ok {
		grpclog.Info("AddRateLimitForMethod ----", name)
		limit := rate.Every(l.opts.interval)
		l.rateLimitMap[name] = rate.NewLimiter(limit, l.opts.burst)
	}
}

func (l *Limit)NewErroringLimiter() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			methodNameByCtx := ctx.Value(define.ReqPatternName)
			pattern :=""
			if methodNameByCtx != nil {
				pattern = methodNameByCtx.(string)
			}
			if len(pattern)>0 {
				rateLimit,ok := l.rateLimitMap[pattern]
				if ok {
					if !rateLimit.Allow() {
						return nil, ratelimit.ErrLimited
					}
				}else{
					l.AddRateLimitForMethod(pattern)
				}
			}
			return next(ctx, request)
		}
	}
}