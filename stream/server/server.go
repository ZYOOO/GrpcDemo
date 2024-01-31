package main

import (
	"EShopDemo/stream/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

const PORT = ":8001"

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) GetStream(req *proto.StreamReqData, serStream proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		err := serStream.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		}) // 非stream情况下得return了之后客户端才能接收到，但是stream情况下Send就能发送数据
		if err != nil {
			return err
		}
		if i > 10 {
			break
		}
		i++
		time.Sleep(time.Second)
	}
	return nil
}

func (s *server) PutStream(clientStream proto.Greeter_PutStreamServer) error {
	for {
		if a, err := clientStream.Recv(); err != nil {
			fmt.Println(err.Error())
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}
func (s *server) AllStream(allStream proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		i := 0
		for {
			data, _ := allStream.Recv()
			fmt.Println("from client:" + data.Data)
			if i > 10 {
				break
			}
			i++
			time.Sleep(time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		i := 0
		for {
			_ = allStream.Send(&proto.StreamResData{Data: "I'm server"})
			if i > 10 {
				break
			}
			i++
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		panic("failed to start server:" + err.Error())
	}
}
