package setting

import (
	"context"
	pb "github.com/LongMarch7/higo/service/admin/setting/pb"
	base "github.com/LongMarch7/higo/service/base"
	endpoint "github.com/go-kit/kit/endpoint"
)

// MakeSayHelloServerEndpoint returns an endpoint that invokes SayHello on the service.
func MakeSayHelloServerEndpoint(s SettingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pb.SayHelloRequest)
		rs, err := s.SayHello(ctx, req.S)
		return pb.SayHelloReply{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// MakeDeleteuserServerEndpoint returns an endpoint that invokes Deleteuser on the service.
func MakeDeleteuserServerEndpoint(s SettingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pb.DeleteuserRequest)
		rs, err := s.Deleteuser(ctx, req.S)
		return pb.DeleteuserReply{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// SayHello implements Service. Primarily useful in a client.
type SayHelloFunc func(ctx context.Context, s []*TestAlias) (rs string, err string)

func SayHelloProxy(e endpoint.Endpoint) SayHelloFunc {
	return func(ctx context.Context, s []*TestAlias) (rs string, err string) {
		request := pb.SayHelloRequest{S: s}
		parameter := base.GrpcClientParameter{
			Method:     "SayHello",
			NewRlyFunc: func() interface{} { return pb.SayHelloReply{} },
			Srv:        "pb.Setting",
		}
		ctx = context.WithValue(ctx, "parameter", parameter)
		response, err := e(ctx, request)
		if err != nil {
			return
		}
		return response.(pb.SayHelloReply).Rs, response.(pb.SayHelloReply).Err
	}
}

// Deleteuser implements Service. Primarily useful in a client.
type DeleteuserFunc func(ctx context.Context, s string) (rs string, err string)

func DeleteuserProxy(e endpoint.Endpoint) DeleteuserFunc {
	return func(ctx context.Context, s string) (rs string, err string) {
		request := pb.DeleteuserRequest{S: s}
		parameter := base.GrpcClientParameter{
			Method:     "Deleteuser",
			NewRlyFunc: func() interface{} { return pb.DeleteuserReply{} },
			Srv:        "pb.Setting",
		}
		ctx = context.WithValue(ctx, "parameter", parameter)
		response, err := e(ctx, request)
		if err != nil {
			return
		}
		return response.(pb.DeleteuserReply).Rs, response.(pb.DeleteuserReply).Err
	}
}
