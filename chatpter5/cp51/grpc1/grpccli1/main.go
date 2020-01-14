package main

import (
	pbh "chatpter5/cp51/grpc1/commpb"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	dial, err := grpc.Dial("127.0.0.1:49971", grpc.WithInsecure())
	if err != nil {
		log.Panicf("did not connect:%v", err)
	}
	defer dial.Close()
	client := pbh.NewGreeterClient(dial)

	name := "world"
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	hello, err := client.SayHello(timeout, &pbh.HelloRequest{
		Name:                 "hello" + name,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	if err != nil {
		log.Panicf("clould not greet:%v", err)
	}
	log.Printf("greeting:%s", hello.Message)
}
