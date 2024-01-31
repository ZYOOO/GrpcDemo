package main

import (
	proto_test "EShopDemo/helloworld/proto"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	req := proto_test.HelloRequest{
		Name: "zhangyong",
	}
	rsp, err := proto.Marshal(&req)
	if err != nil {
		return
	}

	fmt.Println(rsp)
	unrsp := proto_test.HelloRequest{}
	fmt.Println(proto.Unmarshal(rsp, &unrsp), unrsp)

}
