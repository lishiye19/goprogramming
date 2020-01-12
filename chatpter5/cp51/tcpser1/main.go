package main

import (
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:49999")
	if err != nil {
		log.Panicln(err.Error())
	}
	log.Println("ct5 listening at:", addr)
	tcp, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Panicln(err.Error())
	}
	accept, err := tcp.Accept()
	if err != nil {
		log.Panicln(err.Error())

	}

	for {

		recbt := make([]byte, 1024)
		read, err := accept.Read(recbt)
		if err != nil {
			//log.Println(err.Error())
			continue
		}
		if read > 0 {
			log.Println("jieshoudaoxinxi:", recbt)
			accept.Write([]byte("jieshoudaoxinxi"))
		}
	}

}
