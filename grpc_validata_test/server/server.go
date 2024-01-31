package main

import (
	validate_test "EShopDemo/grpc_validata_test/proto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	validate_test.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *validate_test.Person) (*validate_test.Person,
	error) {
	return &validate_test.Person{
		Id:    32,
		Email: "1231@12.com",
	}, nil
}

type Validator interface {
	Validate() error
}

func main() {
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 继续处理请求
		// 如果转换成validate.Person格式的话可以满足person的验证，但是这个拦截器是用于所有的接口的，所以用Validator接口代替
		if r, ok := req.(Validator); ok {
			if err := r.Validate(); err != nil {
				// status, codes  grpc中内置的状态码
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		return handler(ctx, req)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	g := grpc.NewServer(opts...)
	validate_test.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8001")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
