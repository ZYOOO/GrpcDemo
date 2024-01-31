package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"

	grpc_test "EShopDemo/grpc_test/proto"
)

type Server struct {
	grpc_test.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *grpc_test.HelloRequest) (*grpc_test.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get metadata error")
	}
	for key, val := range md {
		fmt.Println(key, ":", val)
	}
	return &grpc_test.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("received a new request")
		res, err := handler(ctx, req)
		fmt.Println("processing complete ")
		return res, err
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	g := grpc.NewServer(opts...)
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
