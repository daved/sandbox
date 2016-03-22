package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/codemodus/parth"
	"github.com/daved/rpctime"
)

type TimeJSON struct {
	Time string `json:"time"`
}

type StatsJSON struct {
	RPCCount uint64 `json:"rpc_count"`
}

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

func (n *node) localHandler(w http.ResponseWriter, r *http.Request) {
	t := &TimeJSON{time.Now().String()}

	b, err := json.Marshal(t)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (n *node) remoteHandler(w http.ResponseWriter, r *http.Request) {
	zoneID, err := parth.SubSegToString(r.URL.Path, "remote")
	if err != nil || zoneID == "" {
		zoneID = "GMT"
	}

	res, err := n.timeServer.Time(zoneID)
	if err != nil {
		if err.Error() == rpctime.ErrZoneNotFound.Error() {
			http.Error(w, "Unprocessable Entity", 422)
			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	t := &TimeJSON{res}

	b, err := json.Marshal(t)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (n *node) statsHandler(w http.ResponseWriter, r *http.Request) {
	res, err := n.timeServer.Stats()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	s := &StatsJSON{res}

	b, err := json.Marshal(s)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (n *node) mux() *http.ServeMux {
	m := http.NewServeMux()

	m.HandleFunc("/api/local", n.localHandler)
	m.HandleFunc("/api/remote/", n.remoteHandler)
	m.HandleFunc("/api/stats", n.statsHandler)

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
