package main

import (
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:49991")
	if err != nil {
		log.Panicln("ResolveUDPAddr error", err.Error())
	}
	udp, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Panicln("DialUDP error", err.Error())
	}
	toUDP, err := udp.Write([]byte("send to ser msg "))
	if err != nil {
		log.Println("WriteToUDP error", err.Error())
	}
	if toUDP > 0 {
		var buf [1024]byte
		read, _ := udp.Read(buf[0:])
		if read > 0 {
			log.Println("Read from ser ", buf)
		}
	}
}
