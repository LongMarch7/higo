package hystrix

import (
	"github.com/LongMarch7/higo/util/define"
	"github.com/afex/hystrix-go/hystrix"
	"context"
	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc/grpclog"
	"sync"
)


type Hystrix struct{
	opts        HystrixOpt
	hystrixMap  map[string] bool
	config      hystrix.CommandConfig
}

var initOpt sync.Once
var hys *Hystrix
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
	initOpt.Do(func() {
		opt := defaultConfig()
		for _, o := range opts {
			o(&opt)
		}
		hys =  &Hystrix{
			opts: opt,
			hystrixMap: make(map[string] bool),
			config: hystrix.CommandConfig{
				Timeout: opt.timeout,
				ErrorPercentThreshold: opt.errorPercentThreshold,
				SleepWindow: opt.sleepWindow,
				MaxConcurrentRequests: opt.maxConcurrentRequests,
				RequestVolumeThreshold: opt.requestVolumeThreshold},
		}
	})
	return hys
}

func (h *Hystrix)Middleware() endpoint.Middleware{
	return h.HystrixMiddleware()
}

func (h *Hystrix)AddHystrixForMethod(name string){
	_,ok := h.hystrixMap[name]
	if !ok{
		grpclog.Info("AddHystrixForMethod ----",name)
		h.hystrixMap[name] =true
		hystrix.ConfigureCommand(name, h.config)
	}
}

func (h *Hystrix)HystrixMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			methodNameByCtx := ctx.Value(define.ReqPatternName)
			pattern :=""
			if methodNameByCtx != nil {
				pattern = methodNameByCtx.(string)
			}
			if len(pattern)>0 {
				_, ok := h.hystrixMap[pattern]
				if ok {
					var resp interface{}
					if err := hystrix.Do(pattern, func() (err error) {
						resp, err = next(ctx, request)
						return err
					}, nil); err != nil {
						return nil, err
					}
					return resp, nil
				}else{
					h.AddHystrixForMethod(pattern)
				}
			}
			return next(ctx, request)
		}
	}
}

