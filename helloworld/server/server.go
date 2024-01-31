package main

import (
	"EShopDemo/helloworld/handler"
	"EShopDemo/helloworld/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		return
	}
	err = server_proxy.RegisterHelloService(&handler.HelloService{})
	if err != nil {
		return
	}
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
