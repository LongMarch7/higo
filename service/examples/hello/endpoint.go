package hello

import (
	"context"
	base "github.com/LongMarch7/higo/service/base"
	pb "github.com/LongMarch7/higo/service/examples/hello/pb"
	endpoint "github.com/go-kit/kit/endpoint"
)

// MakeHelloWorldServerEndpoint returns an endpoint that invokes HelloWorld on the service.
func MakeHelloWorldServerEndpoint(s HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.HelloWorldRequest)
		rs, err := s.HelloWorld(ctx, req.S)
		return &pb.HelloWorldReply{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// HelloWorld implements Service. Primarily useful in a client.
type HelloWorldFunc func(ctx context.Context, s string) (rs string, err string)

func HelloWorldProxy(e endpoint.Endpoint) HelloWorldFunc {
	return func(ctx context.Context, s string) (rs string, err string) {
		request := &pb.HelloWorldRequest{S: s}
		parameter := base.GrpcClientParameter{
			Method:     "HelloWorld",
			NewRlyFunc: func() interface{} { return &pb.HelloWorldReply{} },
			Srv:        "pb.Hello",
		}
		ctx = context.WithValue(ctx, "parameter", parameter)
		response, grpcErr := e(ctx, request)
		if grpcErr != nil {
			err = grpcErr.Error()
			return
		}
		return response.(*pb.HelloWorldReply).Rs, response.(*pb.HelloWorldReply).Err
	}
}
