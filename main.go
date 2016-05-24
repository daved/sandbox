package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/codemodus/sigmon"
	"github.com/codemodus/vitals"
	"github.com/tylerb/graceful"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "systemd")
}

func main() {
	// ignore os signals
	sm := sigmon.New(nil)
	sm.Run()

	// setup and cleanup pid file
	cleanupPID, err := vitals.SetupPIDFile()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer cleanupPID()

	// setup http server
	m := http.NewServeMux()
	m.HandleFunc("/", handler)
	s := &graceful.Server{
		Timeout: time.Second * 3,
		Server: &http.Server{
			Handler: m,
			Addr:    ":12122",
		},
	}

	// handle os signals
	sm.Set(func(*sigmon.SignalMonitor) {
		s.Stop(time.Second * 3)
	})

	// listen and serve
	if err = s.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
