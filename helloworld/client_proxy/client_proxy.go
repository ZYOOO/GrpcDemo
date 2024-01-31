package client_proxy

import (
	"EShopDemo/helloworld/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protocol, address string) HelloServiceStub {
	conn, err := rpc.Dial(protocol, address)
	if err != nil {
		panic("connect errpr")
	}
	return HelloServiceStub{conn}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
