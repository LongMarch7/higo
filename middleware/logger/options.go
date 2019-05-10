package logger


type LoggerOpt struct {
	prefix           string
	methodName       string
}
type LOption func(o *LoggerOpt)


func MethodName(methodName string) LOption {
	return func(o *LoggerOpt) {
		o.methodName = methodName
	}
}

func Prefix(prefix string) LOption {
	return func(o *LoggerOpt) {
		o.prefix = prefix
	}
}

