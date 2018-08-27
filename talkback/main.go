// talkback is a simple demo/intro app for udp communication.
//
//  Available flags:
//
//  --port={port}    set port for udp communication    // ":8089"
//  --ip={ip}        set ip for udp dialing            // "192.168.0.42"
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

type opts struct {
	port     string
	remoteIP string
}

// TODO: validate opts

func main() {
	o := opts{}
	{
		flag.StringVar(&o.port, "port", ":8089", "port for udp communication")
		flag.StringVar(&o.remoteIP, "ip", "192.168.0.42", "ip for udp dialing")
	}
	flag.Parse()

	go func() {
		if err := listen(o.port); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("test")
	}()

	if err := send(o.remoteIP + o.port); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func listen(addr string) error {
	ua, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", ua)
	if err != nil {
		return err
	}

	b := make([]byte, 1024)

	for {
		n, _, err := conn.ReadFromUDP(b)
		if err != nil {
			return err
		}
		s := string(b[:n])

		fmt.Printf("received: %s", s)
	}
}

func send(addr string) error {
	ua, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, ua)
	if err != nil {
		return err
	}

	for {
		r := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		b, err := r.ReadBytes('\n')
		if err != nil {
			return err
		}

		_, err = conn.Write(b)
		if err != nil {
			return err
		}
	}
}
