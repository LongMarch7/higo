package setting

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
)

// SayHelloRequest collects the request parameters for the SayHello method.
type SayHelloRequest struct {
	S string `json:"s"`
}

// SayHelloResponse collects the response parameters for the SayHello method.
type SayHelloResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeSayHelloServerEndpoint returns an endpoint that invokes SayHello on the service.
func MakeSayHelloServerEndpoint(s SettingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SayHelloRequest)
		rs, err := s.SayHello(ctx, req.S)
		return SayHelloResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r SayHelloResponse) Failed() error {
	return r.Err
}

// DeleteuserRequest collects the request parameters for the Deleteuser method.
type DeleteuserRequest struct {
	S string `json:"s"`
}

// DeleteuserResponse collects the response parameters for the Deleteuser method.
type DeleteuserResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDeleteuserServerEndpoint returns an endpoint that invokes Deleteuser on the service.
func MakeDeleteuserServerEndpoint(s SettingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteuserRequest)
		rs, err := s.Deleteuser(ctx, req.S)
		return DeleteuserResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteuserResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SayHello implements Service. Primarily useful in a client.
type SayHelloFunc func(ctx context.Context, s string) (rs string, err error)

func SayHelloProxy(e endpoint.Endpoint) SayHelloFunc {
	return func(ctx context.Context, s string) (rs string, err error) {
		request := SayHelloRequest{S: s}
		ctx = context.WithValue(ctx, "srv", "pb.Setting")
		ctx = context.WithValue(ctx, "method", "SayHello")
		response, err := e(ctx, request)
		if err != nil {
			return
		}
		return response.(SayHelloResponse).Rs, response.(SayHelloResponse).Err
	}
}

// Deleteuser implements Service. Primarily useful in a client.
type DeleteuserFunc func(ctx context.Context, s string) (rs string, err error)

func DeleteuserProxy(e endpoint.Endpoint) DeleteuserFunc {
	return func(ctx context.Context, s string) (rs string, err error) {
		request := DeleteuserRequest{S: s}
		ctx = context.WithValue(ctx, "srv", "pb.Setting")
		ctx = context.WithValue(ctx, "method", "Deleteuser")
		response, err := e(ctx, request)
		if err != nil {
			return
		}
		return response.(DeleteuserResponse).Rs, response.(DeleteuserResponse).Err
	}
}
