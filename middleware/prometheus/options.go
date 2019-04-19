package prometheus


type PrometheusOpt struct {
	fieldKeys []string
	namespace string
	subsystem string
	name      string
	class      int8
	help      string
	buckets []float64
	lvs []string
	count float64
	isSet bool
}
type POption func(o *PrometheusOpt)


func FieldKeys(fieldKeys []string) POption {
	return func(o *PrometheusOpt) {
		o.fieldKeys = fieldKeys
	}
}

func Namespace(namespace string) POption {
	return func(o *PrometheusOpt) {
		o.namespace = namespace
	}
}

func Subsystem(subsystem string) POption {
	return func(o *PrometheusOpt) {
		o.subsystem = subsystem
	}
}

func Name(name string) POption {
	return func(o *PrometheusOpt) {
		o.name = name
	}
}
func GetName(p *Prometheus) string {
	return p.opts.name
}

func Class(class int8) POption {
	return func(o *PrometheusOpt) {
		o.class = class
	}
}

func Help(help string) POption {
	return func(o *PrometheusOpt) {
		o.help = help
	}
}

func Buckets(buckets []float64) POption {
	return func(o *PrometheusOpt) {
		o.buckets = buckets
	}
}

func Lvs(lvs []string) POption {
	return func(o *PrometheusOpt) {
		o.lvs = lvs
	}
}

func Count(count float64) POption {
	return func(o *PrometheusOpt) {
		o.count = count
	}
}

func IsSet(isSet bool) POption {
	return func(o *PrometheusOpt) {
		o.isSet = isSet
	}
}