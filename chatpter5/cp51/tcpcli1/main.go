package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:49999")
	if err != nil {
		log.Panicln(err.Error())
	}
	tcp, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Panicln(err.Error())
	}
	for i := 0; i < 20; i++ {

		write, err := tcp.Write([]byte(fmt.Sprintf("send msg to server [%d]", i)))
		if err != nil {
			log.Panicln(err.Error())
		}
		if write > 0 {
			log.Println(fmt.Sprintf("send msg to server [%d]", i))
		}
		time.Sleep(time.Second * 5)
	}
}
