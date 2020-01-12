package main

//
//import (
//	"log"
//	"net"
//)
//
//func main() {
//	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:49999")
//	if err != nil {
//		log.Panicln(err.Error())
//	}
//	log.Println("ct5 listening at:", addr)
//	tcp, err := net.ListenTCP("tcp", addr)
//	if err != nil {
//		log.Panicln(err.Error())
//	}
//	accept, err := tcp.Accept()
//	if err != nil {
//		log.Panicln(err.Error())
//
//	}
//	for {
//
//		recbt := make([]byte, 1024)
//		accept.Read(recbt)
//	}
//
//}
