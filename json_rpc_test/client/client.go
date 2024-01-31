package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 使用net拨号
	conn, _ := net.Dial("tcp", "localhost:8002")
	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	_ = client.Call("HelloService.Hello", "zhangyong", &reply)
	fmt.Println(reply)
}
