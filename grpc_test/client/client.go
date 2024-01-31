package main

import (
	grpc_test "EShopDemo/grpc_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := grpc_test.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &grpc_test.HelloRequest{Name: "zhangyong"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
