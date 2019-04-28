package test

import (
	pb "github.com/LongMarch7/higo/service/examples/test/pb"
	context "golang.org/x/net/context"
)

func (g *GrpcServer) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloReply, error) {
	_, rep, err := g.SayHelloHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SayHelloReply), nil
}
func (g *GrpcServer) Deleteuser(ctx context.Context, req *pb.DeleteuserRequest) (*pb.DeleteuserReply, error) {
	_, rep, err := g.DeleteuserHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteuserReply), nil
}
func (g *GrpcServer) TestArray(ctx context.Context, req *pb.TestArrayRequest) (*pb.TestArrayReply, error) {
	_, rep, err := g.TestArrayHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.TestArrayReply), nil
}
