package hystrix

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
)


type Hystrix struct{
	opts HystrixOpt
}

func defaultConfig() HystrixOpt{
	return HystrixOpt{
		name:          "/gateway/default/hystrix",
		timeout:                1000,
		maxConcurrentRequests:  100,
		requestVolumeThreshold: 50,
		sleepWindow:            5000,
		errorPercentThreshold:  50,
	}
}

func NewHystrix(opts ...HOption) *Hystrix{
	opt := defaultConfig()
	for _, o := range opts {
		o(&opt)
	}
	return &Hystrix{
		opts: opt,
	}
}

func (h *Hystrix)Middleware() endpoint.Middleware{
	hystrix.ConfigureCommand(h.opts.name, hystrix.CommandConfig{
		Timeout: h.opts.timeout,
		ErrorPercentThreshold: h.opts.errorPercentThreshold,
		SleepWindow: h.opts.sleepWindow,
		MaxConcurrentRequests: h.opts.maxConcurrentRequests,
		RequestVolumeThreshold: h.opts.requestVolumeThreshold,
	})
	return circuitbreaker.Hystrix(h.opts.name)
}

