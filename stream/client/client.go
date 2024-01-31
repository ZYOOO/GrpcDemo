package main

import (
	"EShopDemo/stream/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	res, err := c.GetStream(context.Background(), &proto.StreamReqData{Data: "yong"})
	if err != nil {
		panic(err)
	}
	for {
		a, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a)
	}

	putS, _ := c.PutStream(context.Background())
	i := 0
	for {
		putS.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("%d yong", i),
		})
		if i > 10 {
			break
		}
		i++
		time.Sleep(time.Second)
	}
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		i := 0
		for {
			data, _ := allStr.Recv()
			fmt.Println("from server:" + data.Data)
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
			_ = allStr.Send(&proto.StreamReqData{Data: "I'm client"})
			if i > 10 {
				break
			}
			i++
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
