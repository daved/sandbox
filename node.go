package main

import (
	"net/http"
	"time"

	"github.com/daved/rpctime"
)

type node struct {
	timeServer  *rpctime.Client
	multiplexer *http.ServeMux
}

func newNode() (*node, error) {
	ts, err := rpctime.NewClient(":19876", time.Second*6)
	if err != nil {
		return nil, err
	}

	n := &node{}
	n.multiplexer = n.mux()
	n.timeServer = ts

	return n, nil
}

func (h *node) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.multiplexer.ServeHTTP(w, r)
}

func (n *node) mux() *http.ServeMux {
	m := http.NewServeMux()

	m.HandleFunc("/api/local", n.localHandler)
	m.HandleFunc("/api/remote/", n.remoteHandler)
	m.HandleFunc("/api/stats", n.statsHandler)

	return m
}
