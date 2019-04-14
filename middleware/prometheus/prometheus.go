package prometheus

import (
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"context"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"time"
)

const (
	Counter_TYPE      int8 = 0
	Summary_TYPE      int8 = 1
	Gauge_TYPE        int8 = 2
	Histogram_TYPE    int8 = 3
)

type PrometheusFunc func(time.Time,error)

type PrometheusEndpoint struct{
	prometheusFunc []PrometheusFunc
}

type Prometheus struct{
	opts               PrometheusOpt
	prometheusEndpoint *PrometheusEndpoint
}

func defaultConfig() PrometheusOpt{
	return PrometheusOpt{
		namespace: "default_space",
		subsystem: "default_sub",
		name: "default_name",
		help: "default help.",
		fieldKeys:  []string{"method", "error"},
		count: 1,
		lvs: []string{"method", "default", "error"},
	}
}

func NewPrometheus(opts ...POption) *Prometheus{
	opt := defaultConfig()
	for _, o := range opts {
		o(&opt)
	}
	return &Prometheus{
		opts: opt,
		prometheusEndpoint: new(PrometheusEndpoint),
	}
}

func (p *Prometheus)Middleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time){
				funcs := p.prometheusEndpoint
				for _, value := range funcs.prometheusFunc{
					value(begin, err)
				}
			}(time.Now())
			return next(ctx, request)
		}
	}
}

func (p *Prometheus)AddObj(opts ...POption) {
	for _, o := range opts {
		o(&p.opts)
	}
	var prometheusFunc PrometheusFunc
	switch p.opts.class{
	case Counter_TYPE:
		requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: p.opts.namespace,
			Subsystem: p.opts.subsystem,
			Name:      p.opts.name,
			Help:      p.opts.help,
		}, p.opts.fieldKeys)
		prometheusFunc = p.PrometheusCounterEndpoint(requestCount)
	case Summary_TYPE:
		requestSummary := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: p.opts.namespace,
			Subsystem: p.opts.subsystem,
			Name:      p.opts.name,
			Help:      p.opts.help,
		}, p.opts.fieldKeys)
		prometheusFunc = p.PrometheusSummaryEndpoint(requestSummary)
	case Gauge_TYPE:
		requestGauge := kitprometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: p.opts.namespace,
			Subsystem: p.opts.subsystem,
			Name:      p.opts.name,
			Help:      p.opts.help,
		}, p.opts.fieldKeys)
		prometheusFunc = p.PrometheusGaugeEndpoint(requestGauge)
	case Histogram_TYPE:
		requestHistogram := kitprometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: p.opts.namespace,
			Subsystem: p.opts.subsystem,
			Name:      p.opts.name,
			Help:      p.opts.help,
			Buckets: p.opts.buckets,
		}, p.opts.fieldKeys)
		prometheusFunc = p.PrometheusHistogramEndpoint(requestHistogram)
	}
	p.prometheusEndpoint.prometheusFunc = append(p.prometheusEndpoint.prometheusFunc, prometheusFunc)
}

func (p *Prometheus)PrometheusCounterEndpoint(counter *kitprometheus.Counter) PrometheusFunc {
	return func(begin time.Time, err error)  {
		lvs := append(p.opts.lvs,fmt.Sprint(err != nil))
		counter.With(lvs...).Add(p.opts.count)
	}
}

func (p *Prometheus)PrometheusSummaryEndpoint(summary *kitprometheus.Summary) PrometheusFunc {
	return func(begin time.Time, err error)  {
		lvs := append(p.opts.lvs,fmt.Sprint(err != nil))
		summary.With(lvs...).Observe(time.Since(begin).Seconds())
	}
}

func (p *Prometheus)PrometheusHistogramEndpoint(histogram *kitprometheus.Histogram) PrometheusFunc {
	return func(begin time.Time, err error)  {
		lvs := append(p.opts.lvs,fmt.Sprint(err != nil))
		histogram.With(lvs...).Observe(time.Since(begin).Seconds())
	}
}

func (p *Prometheus)PrometheusGaugeEndpoint(gauge *kitprometheus.Gauge) PrometheusFunc {
	return func(begin time.Time, err error)  {
		lvs := append(p.opts.lvs,fmt.Sprint(err != nil))
		if p.opts.isSet {
			gauge.With(lvs...).Set(p.opts.count)
		}else{
			gauge.With(lvs...).Add(p.opts.count)
		}
	}
}
