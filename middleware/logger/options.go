package logger


type LoggerOpt struct {
	methodName           string
}
type LOption func(o *LoggerOpt)


func MethodName(methodName string) LOption {
	return func(o *LoggerOpt) {
		o.methodName = methodName
	}
}

