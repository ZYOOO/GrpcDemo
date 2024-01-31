package main

import (
	grpc_test "EShopDemo/grpc_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
