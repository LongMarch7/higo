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
		rs, err := s.HtmlCall(ctx, req.Method, req.Pattern)
		return &pb.HtmlCallReply{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// MakeApiCallServerEndpoint returns an endpoint that invokes ApiCall on the service.
func MakeApiCallServerEndpoint(s WebService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ApiCallRequest)
		rs, err := s.ApiCall(ctx, req.Method, req.Pattern)
		return &pb.ApiCallReply{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// HtmlCall implements Service. Primarily useful in a client.
type HtmlCallFunc func(ctx context.Context, method string, pattern string) (rs string, err string)

func HtmlCallProxy(e endpoint.Endpoint) HtmlCallFunc {
	return func(ctx context.Context, method string, pattern string) (rs string, err string) {
		request := &pb.HtmlCallRequest{
			Method:  method,
			Pattern: pattern,
		}
		parameter := base.GrpcClientParameter{
			Method:     "HtmlCall",
			NewRlyFunc: func() interface{} { return &pb.HtmlCallReply{} },
			Srv:        "pb.Web",
		}
		ctx = context.WithValue(ctx, "parameter", parameter)
		response, grpcErr := e(ctx, request)
		if grpcErr != nil {
			err = grpcErr.Error()
			return
		}
		return response.(*pb.HtmlCallReply).Rs, response.(*pb.HtmlCallReply).Err
	}
}

// ApiCall implements Service. Primarily useful in a client.
type ApiCallFunc func(ctx context.Context, method string, pattern string) (rs string, err string)

func ApiCallProxy(e endpoint.Endpoint) ApiCallFunc {
	return func(ctx context.Context, method string, pattern string) (rs string, err string) {
		request := &pb.ApiCallRequest{
			Method:  method,
			Pattern: pattern,
		}
		parameter := base.GrpcClientParameter{
			Method:     "ApiCall",
			NewRlyFunc: func() interface{} { return &pb.ApiCallReply{} },
			Srv:        "pb.Web",
		}
		ctx = context.WithValue(ctx, "parameter", parameter)
		response, grpcErr := e(ctx, request)
		if grpcErr != nil {
			err = grpcErr.Error()
			return
		}
		return response.(*pb.ApiCallReply).Rs, response.(*pb.ApiCallReply).Err
	}
}
