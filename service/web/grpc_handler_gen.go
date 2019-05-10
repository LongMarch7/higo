// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package web

import (
	pb "github.com/LongMarch7/higo/service/web/pb"
	define "github.com/LongMarch7/higo/util/define"
	grpc "github.com/go-kit/kit/transport/grpc"
	grpc1 "google.golang.org/grpc"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type GrpcServer struct {
	HtmlCallHandler grpc.Handler
	ApiCallHandler  grpc.Handler
}

func MakeRegisteFunc(srv pb.WebServer) define.GrpcRegister {
	return func(s *grpc1.Server) {
		pb.RegisterWebServer(s, srv)
	}
}