package ratelimit

import (
	"time"
)

type ratelimitOpt struct {
	interval time.Duration
	burst int
}
type ROption func(o *ratelimitOpt)


func Interval(interval time.Duration) ROption {
	return func(o *ratelimitOpt) {
		o.interval = interval
	}
}

func Burst(burst int) ROption {
	return func(o *ratelimitOpt) {
		o.burst = burst
	}
}
