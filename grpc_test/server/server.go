package main

import (
	"context"
	"google.golang.org/grpc"
	"net"

	grpc_test "EShopDemo/grpc_test/proto"
)

type Server struct {
	grpc_test.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *grpc_test.HelloRequest) (*grpc_test.HelloReply, error) {
	return &grpc_test.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	grpc_test.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8001")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
