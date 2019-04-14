package ratelimit

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
	"time"
)

type Limit struct{
	opts ratelimitOpt
}

func defaultConfig() ratelimitOpt{
	return ratelimitOpt{
		rateLimit: rate.Every(time.Millisecond * 10),
		burst: 100,
	}
}

func NewLimiter(opts ...ROption) *Limit{
	opt := defaultConfig()
	for _, o := range opts {
		o(&opt)
	}
	return &Limit{
		opts: opt,
	}
}

func (l *Limit)Middleware() endpoint.Middleware {
	return ratelimit.NewErroringLimiter(rate.NewLimiter(l.opts.rateLimit, l.opts.burst))
}