package ratelimit

import "golang.org/x/time/rate"

type ratelimitOpt struct {
	rateLimit rate.Limit
	burst int
}
type ROption func(o *ratelimitOpt)


func RateLimit(rateLimit rate.Limit) ROption {
	return func(o *ratelimitOpt) {
		o.rateLimit = rateLimit
	}
}

func Burst(burst int) ROption {
	return func(o *ratelimitOpt) {
		o.burst = burst
	}
}
