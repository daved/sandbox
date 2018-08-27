package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":33233")
	if err != nil {
		log.Fatalln(err)
	}

	l, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		b := make([]byte, 1024)
		n, err := l.Read(b)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Print(string(b[:n]))
	}
}
