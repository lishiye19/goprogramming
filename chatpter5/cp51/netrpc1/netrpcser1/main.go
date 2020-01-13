package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Arith struct {
}

type ArithRequest struct {
	A int
	B int
}

type ArithResponse struct {
	Pro int
	Quo int
	Rem int
}

func main() {
	///*****************net/rpc begin*******************///
	/***
	rpc.Register(new(Arith))
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", "127.0.0.1:49981")
	if err != nil {
		log.Panicln("net.Listen err", err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", "开始链接")
	log.Println("开始监听", listen.Addr())
	http.Serve(listen, nil)
	///*****************net/rpc end*******************/ //
	///**********net/rpc/jsonrpc begin*******///
	///*******//
	rpc.Register(new(Arith))
	listen, err := net.Listen("tcp", "127.0.0.1:49982")
	if err != nil {
		log.Panicln("开始监听:", listen.Addr())
	}
	log.Printf("开始监听:%s\n", listen.Addr())
	for {
		accept, err := listen.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			log.Printf("建立链接%s\n", conn.RemoteAddr())
			jsonrpc.ServeConn(conn)
		}(accept)
	}
	///*****************net/rpc/jsonrpc end*******************///
}

func (at *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

func (at *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("req.B 为0,不符合规则")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}
