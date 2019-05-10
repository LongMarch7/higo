package web

import (
	pb "github.com/LongMarch7/higo/service/web/pb"
	context "golang.org/x/net/context"
)

func (g *GrpcServer) HtmlCall(ctx context.Context, req *pb.HtmlCallRequest) (*pb.HtmlCallReply, error) {
	_, rep, err := g.HtmlCallHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HtmlCallReply), nil
}
func (g *GrpcServer) ApiCall(ctx context.Context, req *pb.ApiCallRequest) (*pb.ApiCallReply, error) {
	_, rep, err := g.ApiCallHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ApiCallReply), nil
}
