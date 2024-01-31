package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, _ := rpc.Dial("tcp", "localhost:8002")
	var reply string
	_ = client.Call("HelloService.Hello", "zhangyong", &reply)
	fmt.Println(reply)
}
