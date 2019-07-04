package prometheus

import (
	"fmt"
	"github.com/LongMarch7/higo/util/define"
	"github.com/LongMarch7/higo/util/global"
	"github.com/go-kit/kit/endpoint"
	"context"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"sync"
	"time"
)

const (
	Counter_TYPE      int8 = 0
	Summary_TYPE      int8 = 1
	Gauge_TYPE        int8 = 2
	Histogram_TYPE    int8 = 3
)

type Prometheus struct{
	opts               PrometheusOpt
	Counter            map[string]*kitprometheus.Counter
	Summary            map[string]*kitprometheus.Summary
	Histogram          map[string]*kitprometheus.Histogram
	Gauge              map[string]*kitprometheus.Gauge
}

func defaultConfig() PrometheusOpt{
	return PrometheusOpt{
		namespace: "higo",
		subsystem: "default_sub",
		name: "default_name",
		help: "default help.",
		fieldKeys:  []string{"method", "error"},
		count: 1,
		lvs: []string{"method", "default", "error"},
	}
}

var initOpt sync.Once
var prometheus *Prometheus

func NewPrometheus(opts ...POption) *Prometheus{
	initOpt.Do(func() {
		opt := defaultConfig()
		for _, o := range opts {
			o(&opt)
		}
		prometheus = &Prometheus{
			opts:      opt,
			Counter:   make(map[string]*kitprometheus.Counter),
			Summary:   make(map[string]*kitprometheus.Summary),
			Histogram: make(map[string]*kitprometheus.Histogram),
			Gauge:     make(map[string]*kitprometheus.Gauge),
		}
	})
	return prometheus
}

func (p *Prometheus)Middleware(opts ...POption) endpoint.Middleware {
	for _, o := range opts {
		o(&p.opts)
	}
	p.AddObj()
	switch p.opts.class{
	case Counter_TYPE:
		if counter, ok := p.Counter[p.opts.subsystem + p.opts.name]; ok {
			return p.PrometheusCounterEndpoint(counter)
		}
	case Summary_TYPE:
		if summary, ok := p.Summary[p.opts.subsystem + p.opts.name]; ok {
			return p.PrometheusSummaryEndpoint(summary)
		}
	case Gauge_TYPE:
		if gauge, ok := p.Gauge[p.opts.subsystem + p.opts.name]; ok {
			return p.PrometheusGaugeEndpoint(gauge)
		}
	case Histogram_TYPE:
		if histogram, ok := p.Histogram[p.opts.subsystem + p.opts.name]; ok {
			return p.PrometheusHistogramEndpoint(histogram)
		}
	}
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return next(ctx, request)
		}
	}
}

func (p *Prometheus)AddObj() {
	switch p.opts.class{
	case Counter_TYPE:
		if _, ok := p.Counter[p.opts.subsystem + p.opts.name]; !ok {
			requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
				Namespace: p.opts.namespace,
				Subsystem: p.opts.subsystem,
				Name:      p.opts.name,
				Help:      p.opts.help,
			}, p.opts.fieldKeys)
			p.Counter[p.opts.subsystem+p.opts.name] = requestCount
		}
	case Summary_TYPE:
		if _, ok := p.Summary[p.opts.subsystem + p.opts.name]; !ok {
			requestSummary := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: p.opts.namespace,
			Subsystem: p.opts.subsystem,
			Name:      p.opts.name,
			Help:      p.opts.help,
		}, p.opts.fieldKeys)
			p.Summary[p.opts.subsystem+p.opts.name] = requestSummary
		}
	case Gauge_TYPE:
		if _, ok := p.Gauge[p.opts.subsystem + p.opts.name]; !ok {
			requestGauge := kitprometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
				Namespace: p.opts.namespace,
				Subsystem: p.opts.subsystem,
				Name:      p.opts.name,
				Help:      p.opts.help,
			}, p.opts.fieldKeys)
			p.Gauge[p.opts.subsystem+p.opts.name] = requestGauge
		}
	case Histogram_TYPE:
		if _, ok := p.Histogram[p.opts.subsystem + p.opts.name]; !ok {
			requestHistogram := kitprometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
				Namespace: p.opts.namespace,
				Subsystem: p.opts.subsystem,
				Name:      p.opts.name,
				Help:      p.opts.help,
				Buckets: p.opts.buckets,
			}, p.opts.fieldKeys)
			p.Histogram[p.opts.subsystem+p.opts.name] = requestHistogram
		}
	}
}

func (p *Prometheus)PrometheusCounterEndpoint(counter *kitprometheus.Counter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func() {
				lvs := LvsByContext(ctx,err)
				if len(lvs) == 0{
					lvs = append(p.opts.lvs, fmt.Sprint(err != nil))
				}
				counter.With(lvs...).Add(p.opts.count)
			}()
			return next(ctx, request)
		}
	}
}

func (p *Prometheus)PrometheusSummaryEndpoint(summary *kitprometheus.Summary) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				lvs := LvsByContext(ctx,err)
				if len(lvs) == 0{
					lvs = append(p.opts.lvs, fmt.Sprint(err != nil))
				}
				summary.With(lvs...).Observe(time.Since(begin).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}

func (p *Prometheus)PrometheusHistogramEndpoint(histogram *kitprometheus.Histogram) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				lvs := LvsByContext(ctx,err)
				if len(lvs) == 0{
					lvs = append(p.opts.lvs, fmt.Sprint(err != nil))
				}
				histogram.With(lvs...).Observe(time.Since(begin).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}

func (p *Prometheus)PrometheusGaugeEndpoint(gauge *kitprometheus.Gauge) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func() {
				lvs := LvsByContext(ctx,err)
				if len(lvs) == 0{
					lvs = append(p.opts.lvs, fmt.Sprint(err != nil))
				}
				if p.opts.isSet {
					gauge.With(lvs...).Set(p.opts.count)
				}else{
					gauge.With(lvs...).Add(p.opts.count)
				}
			}()
			return next(ctx, request)
		}
	}
}

func LvsByContext(ctx context.Context,err error) []string{
	var lvs []string
	if global.AppMode == define.SvrMode {
		methodNameByCtx := ctx.Value(define.ReqPatternName)
		if methodNameByCtx != nil {
			lvs =append(lvs, "method", methodNameByCtx.(string), "error",fmt.Sprint(err != nil))
		}
	}
	return lvs
}
