package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		return
	}
	err = rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		return
	}
	for {
		conn, _ := listener.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
