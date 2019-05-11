package web

import (
	"context"

	base "github.com/LongMarch7/higo/service/base"
	pb "github.com/LongMarch7/higo/service/web/pb"
	endpoint "github.com/go-kit/kit/endpoint"
)

// MakeHtmlCallServerEndpoint returns an endpoint that invokes HtmlCall on the service.
func MakeHtmlCallServerEndpoint(s WebService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.HtmlCallRequest)
		rs, err := s.HtmlCall(ctx, req.Pattern)
		return &pb.HtmlCallReply{Rs: rs}, err
	}
}

// MakeApiCallServerEndpoint returns an endpoint that invokes ApiCall on the service.
func MakeApiCallServerEndpoint(s WebService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ApiCallRequest)
		rs, err := s.ApiCall(ctx, req.Pattern)
		return &pb.ApiCallReply{Rs: rs}, err
	}
}

// HtmlCall implements Service. Primarily useful in a client.
type HtmlCallFunc func(ctx context.Context, pattern string) (rs string, err error)

func HtmlCallProxy(e endpoint.Endpoint) HtmlCallFunc {
	return func(ctx context.Context, pattern string) (rs string, err error) {
		request := &pb.HtmlCallRequest{Pattern: pattern}
		parameter := base.GrpcClientParameter{
			Method:     "HtmlCall",
			NewRlyFunc: func() interface{} { return &pb.HtmlCallReply{} },
			Srv:        "pb.Web",
		}
		ctx = context.WithValue(ctx, "parameter", parameter)
		response, grpcErr := e(ctx, request)
		if grpcErr != nil {
			err = grpcErr
			return
		}
		return response.(*pb.HtmlCallReply).Rs, nil
	}
}

// ApiCall implements Service. Primarily useful in a client.
type ApiCallFunc func(ctx context.Context, pattern string) (rs string, err error)

func ApiCallProxy(e endpoint.Endpoint) ApiCallFunc {
	return func(ctx context.Context, pattern string) (rs string, err error) {
		request := &pb.ApiCallRequest{Pattern: pattern}
		parameter := base.GrpcClientParameter{
			Method:     "ApiCall",
			NewRlyFunc: func() interface{} { return &pb.ApiCallReply{} },
			Srv:        "pb.Web",
		}
		ctx = context.WithValue(ctx, "parameter", parameter)
		response, grpcErr := e(ctx, request)
		if grpcErr != nil {
			err = grpcErr
			return
		}
		return response.(*pb.ApiCallReply).Rs, nil
	}
}
