package hystrix


type HystrixOpt struct {
	name                   string
	timeout                int
	maxConcurrentRequests  int
	requestVolumeThreshold int
	sleepWindow            int
	errorPercentThreshold  int
}
type HOption func(o *HystrixOpt)


func Name(name string) HOption {
	return func(o *HystrixOpt) {
		o.name = name
	}
}

func Timeout(timeout int) HOption {
	return func(o *HystrixOpt) {
		o.timeout = timeout
	}
}

func MaxConcurrentRequests(maxConcurrentRequests int) HOption {
	return func(o *HystrixOpt) {
		o.maxConcurrentRequests = maxConcurrentRequests
	}
}

func RequestVolumeThreshold(requestVolumeThreshold int) HOption {
	return func(o *HystrixOpt) {
		o.requestVolumeThreshold = requestVolumeThreshold
	}
}

func SleepWindow(sleepWindow int) HOption {
	return func(o *HystrixOpt) {
		o.sleepWindow = sleepWindow
	}
}

func ErrorPercentThreshold(errorPercentThreshold int) HOption {
	return func(o *HystrixOpt) {
		o.errorPercentThreshold = errorPercentThreshold
	}
}