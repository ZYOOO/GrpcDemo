package main

import (
	validate_test "EShopDemo/grpc_validata_test/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption

	//opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:8001", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := validate_test.NewGreeterClient(conn)
	rsp, err := c.SayHello(context.Background(), &validate_test.Person{
		Id:     1000,
		Email:  "zyooooo@126.com",
		Mobile: "15986528666",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", rsp.Email)
}
