package test

import (
	"context"

	base "github.com/LongMarch7/higo/service/base"
	pb "github.com/LongMarch7/higo/service/examples/test/pb"
	endpoint "github.com/go-kit/kit/endpoint"
)

// MakeSayHelloServerEndpoint returns an endpoint that invokes SayHello on the service.
func MakeSayHelloServerEndpoint(s TestService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.SayHelloRequest)
		rs, err := s.SayHello(ctx, req.S)
		return &pb.SayHelloReply{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// MakeDeleteuserServerEndpoint returns an endpoint that invokes Deleteuser on the service.
func MakeDeleteuserServerEndpoint(s TestService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.DeleteuserRequest)
		rs, err := s.Deleteuser(ctx, req.S)
		return &pb.DeleteuserReply{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// MakeTestArrayServerEndpoint returns an endpoint that invokes TestArray on the service.
func MakeTestArrayServerEndpoint(s TestService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.TestArrayRequest)
		rs, err := s.TestArray(ctx, req.S)
		return &pb.TestArrayReply{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// SayHello implements Service. Primarily useful in a client.
type SayHelloFunc func(ctx context.Context, s *TestStrucAlias) (rs string, err string)

func SayHelloProxy(e endpoint.Endpoint) SayHelloFunc {
	return func(ctx context.Context, s *TestStrucAlias) (rs string, err string) {
		request := &pb.SayHelloRequest{S: s}
		parameter := base.GrpcClientParameter{
			Method:     "SayHello",
			NewRlyFunc: func() interface{} { return &pb.SayHelloReply{} },
			Srv:        "pb.Test",
		}
		ctx = context.WithValue(ctx, "parameter", parameter)
		response, grpcErr := e(ctx, request)
		if grpcErr != nil {
			err = grpcErr.Error()
			return
		}
		return response.(*pb.SayHelloReply).Rs, response.(*pb.SayHelloReply).Err
	}
}

// Deleteuser implements Service. Primarily useful in a client.
type DeleteuserFunc func(ctx context.Context, s string) (rs string, err string)

func DeleteuserProxy(e endpoint.Endpoint) DeleteuserFunc {
	return func(ctx context.Context, s string) (rs string, err string) {
		request := &pb.DeleteuserRequest{S: s}
		parameter := base.GrpcClientParameter{
			Method:     "Deleteuser",
			NewRlyFunc: func() interface{} { return &pb.DeleteuserReply{} },
			Srv:        "pb.Test",
		}
		ctx = context.WithValue(ctx, "parameter", parameter)
		response, grpcErr := e(ctx, request)
		if grpcErr != nil {
			err = grpcErr.Error()
			return
		}
		return response.(*pb.DeleteuserReply).Rs, response.(*pb.DeleteuserReply).Err
	}
}

// TestArray implements Service. Primarily useful in a client.
type TestArrayFunc func(ctx context.Context, s []*TestStrucAlias) (rs string, err string)

func TestArrayProxy(e endpoint.Endpoint) TestArrayFunc {
	return func(ctx context.Context, s []*TestStrucAlias) (rs string, err string) {
		request := &pb.TestArrayRequest{S: s}
		parameter := base.GrpcClientParameter{
			Method:     "TestArray",
			NewRlyFunc: func() interface{} { return &pb.TestArrayReply{} },
			Srv:        "pb.Test",
		}
		ctx = context.WithValue(ctx, "parameter", parameter)
		response, grpcErr := e(ctx, request)
		if grpcErr != nil {
			err = grpcErr.Error()
			return
		}
		return response.(*pb.TestArrayReply).Rs, response.(*pb.TestArrayReply).Err
	}
}
