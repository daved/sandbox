package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var httpPort, rpcPort string

	flag.StringVar(&httpPort, "http-port", ":29876", "port on which to handle http requests")
	flag.StringVar(&rpcPort, "rpc-port", ":19876", "port on which to make rpc requests")

	flag.Parse()

	n, err := newNode(rpcPort)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	http.ListenAndServe(httpPort, n)
}
