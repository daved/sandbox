package main

import (
	"flag"
	"log"
	"net"
	"net/rpc"

	"github.com/daved/rpctime"
)

func main() {
	var port string

	flag.StringVar(&port, "port", ":19876", "port on which to handle rpc requests")

	flag.Parse()

	rpc.Register(rpctime.NewRPC())

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	rpc.Accept(l)
}
