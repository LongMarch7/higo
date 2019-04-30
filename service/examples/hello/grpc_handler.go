package hello

import (
	pb "github.com/LongMarch7/higo/service/examples/hello/pb"
	context "golang.org/x/net/context"
)

func (g *GrpcServer) HelloWorld(ctx context.Context, req *pb.HelloWorldRequest) (*pb.HelloWorldReply, error) {
	_, rep, err := g.HelloWorldHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HelloWorldReply), nil
}
