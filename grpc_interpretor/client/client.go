package main

import (
	grpc_test "EShopDemo/grpc_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"time"
)

func main() {
	var interceptor grpc.UnaryClientInterceptor
	interceptor = func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("expend: %s\n", time.Since(start))
		return err
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))

	conn, err := grpc.Dial("127.0.0.1:8001", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := grpc_test.NewGreeterClient(conn)
	md := metadata.New(map[string]string{
		"name":     "zhangyong",
		"password": "zhangyong123",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.SayHello(ctx, &grpc_test.HelloRequest{Name: "zhangyong"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
