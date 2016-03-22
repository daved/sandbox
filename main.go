package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codemodus/parth"
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

func (n *node) localTimeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().String()

	w.Write([]byte(t))
}

func (n *node) remoteTimeHandler(w http.ResponseWriter, r *http.Request) {
	zoneID, err := parth.SubSegToString(r.URL.Path, "zone")
	if err != nil || zoneID == "" {
		zoneID = "GMT"
	}

	t, err := n.timeServer.Time(zoneID)
	if err != nil {
		if err.Error() == rpctime.ErrZoneNotFound.Error() {
			http.Error(w, "Unprocessable Entity", 422)
			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(t))
}

func (n *node) statsHandler(w http.ResponseWriter, r *http.Request) {
	ct, err := n.timeServer.Stats()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(strconv.FormatUint(ct, 10)))
}

func (n *node) mux() *http.ServeMux {
	m := http.NewServeMux()

	m.HandleFunc("/time/local", n.localTimeHandler)
	m.HandleFunc("/time/remote", n.remoteTimeHandler)
	m.HandleFunc("/time/remote/zone/", n.remoteTimeHandler)
	m.HandleFunc("/time/stats", n.statsHandler)

	return m
}

func main() {
	n, err := newNode()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	http.ListenAndServe(":29876", n)
}
