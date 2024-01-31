package main

import (
	"EShopDemo/helloworld/client_proxy"
	"fmt"
)

func main() {
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:8002")
	var reply string
	err := client.Hello("zhangyong", &reply)
	if err != nil {
		panic("call failed")
	}
	fmt.Println(reply)
}
