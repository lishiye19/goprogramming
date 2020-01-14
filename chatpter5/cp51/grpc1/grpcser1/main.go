package main

import (
	pbh "chatpter5/cp51/grpc1/commpb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Serverh struct {
}

func (sh *Serverh) SayHello(ctx context.Context, in *pbh.HelloRequest) (*pbh.HelloReply, error) {
	return &pbh.HelloReply{
		Message:              "Hello " + in.Name,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:49971")
	if err != nil {
		log.Panicf("failed to listen:%v", err)
	}
	server := grpc.NewServer()
	pbh.RegisterGreeterServer(server, &Serverh{})
	reflection.Register(server)
	err = server.Serve(lis)
	if err != nil {
		log.Panicf("failed to serve:%v", err)
	}
}
