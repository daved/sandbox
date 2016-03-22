package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/daved/rpctime"
)

func main() {
	rpc.Register(rpctime.NewRPC())

	l, err := net.Listen("tcp", ":19876")
	if err != nil {
		log.Fatalln(err)
	}

	rpc.Accept(l)
}
