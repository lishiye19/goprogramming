package main

import (
	"log"
	"net"
	"time"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:49991")
	if err != nil {
		log.Panicln("ResolveUDPAddr error", err.Error())
	}
	udp, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Panicln("ListenUDP error", err.Error())
	}

	log.Println("ListenUDP at ", addr)
	for {
		var buf [1024]byte
		from, a, err := udp.ReadFromUDP(buf[0:])
		if err != nil {
			log.Println("ReadFrom error", err.Error())
		}
		if from > 0 {
			log.Println("ReadFrom addr:", a, "msg:", buf)
			write, _ := udp.WriteToUDP([]byte("already recv msg!"), a)
			if write > 0 {
				log.Println("Write to cli len", write, "cli addr:", a)
			}

		}
		time.Sleep(time.Second * 1)
	}

}
