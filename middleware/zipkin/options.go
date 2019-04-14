package zipkin

type zipkinOpt struct {
	name       string
	url        string
	hostPort   string
	debug      bool
	methodName string
}
type ZOption func(o *zipkinOpt)

func Name(name string) ZOption {
	return func(o *zipkinOpt) {
		o.name = name
	}
}

func Url(url string) ZOption {
	return func(o *zipkinOpt) {
		o.url = url
	}
}

func HostPort(hostPort string) ZOption {
	return func(o *zipkinOpt) {
		o.hostPort = hostPort
	}
}

func Debug(debug bool) ZOption {
	return func(o *zipkinOpt) {
		o.debug = debug
	}
}

func MethodName(methodName string) ZOption {
	return func(o *zipkinOpt) {
		o.methodName = methodName
	}
}

