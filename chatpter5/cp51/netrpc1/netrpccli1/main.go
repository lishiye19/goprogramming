package main

import (
	"log"
	"net/rpc/jsonrpc"
)

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
	//////////////*******netrpc cli begin******//////////////
	/******
	http, err := rpc.DialHTTP("tcp", "127.0.0.1:49981")
	if err != nil {
		log.Panicln("rpc.DialHTTP err", err)
	}
	var arq = ArithRequest{11, 2}
	var arp = new(ArithResponse)
	err = http.Call("Arith.Multiply", arq, arp)
	if err != nil {
		log.Println("http.Call[Arith.Multiply] err", err)
	}
	log.Printf("%d * %d=%d\n", arq.A, arq.B, arp.Pro)
	err = http.Call("Arith.Divide", arq, arp)
	if err != nil {
		log.Println("http.Call[Arith.Divide] err", err)
	}
	log.Printf("%d / %d 商[%d] 余[%d]\n", arq.A, arq.B, arp.Quo, arp.Rem)
	///*********netrpc cli end*************/ //
	//////////*************net/rpc/jsonrpc begin**************//////////
	dial, err := jsonrpc.Dial("tcp", "127.0.0.1:49982")
	if err != nil {
		log.Panicln("jsonrpc.Dial err", err)
	}
	arq := ArithRequest{11, 3}
	var arp = new(ArithResponse)
	err = dial.Call("Arith.Multiply", arq, arp)
	if err != nil {
		log.Println("http.Call[Arith.Multiply] err", err)
	}
	log.Printf("%d * %d=%d\n", arq.A, arq.B, arp.Pro)
	err = dial.Call("Arith.Divide", arq, arp)
	if err != nil {
		log.Println("http.Call[Arith.Divide] err", err)
	}
	log.Printf("%d / %d 商[%d] 余[%d]\n", arq.A, arq.B, arp.Quo, arp.Rem)
	//////////*************net/rpc/jsonrpc end**************//////////
}
