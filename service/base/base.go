package base

type GrpcClientParameter struct{
	Srv           string
	Method        string
	NewRlyFunc    func() interface{}
}
