package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/codemodus/parth"
	"github.com/daved/rpctime"
)

type TimeJSON struct {
	Time time.Time `json:"time"`
}

type StatsJSON struct {
	RPCCount uint64 `json:"rpc_count"`
}

func (n *node) localHandler(w http.ResponseWriter, r *http.Request) {
	t := &TimeJSON{time.Now()}

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
