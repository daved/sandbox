package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/codemodus/sigmon"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s [%s] %s\n", time.Now().Format(time.RFC3339), r.Method, r.URL.Path)
	fmt.Fprintln(w, "systemd")
}

func main() {
	fmt.Println("hello")

	// setup and parse flags
	port := ":12121"
	flag.StringVar(
		&port, "port", port,
		"port to listen on for http requests",
	)
	flag.Parse()

	// ignore os signals
	sm := sigmon.New(nil)
	sm.Run()

	// setup http server
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	s := &http.Server{
		Handler: mux,
		Addr:    port,
	}

	// handle os signals
	sm.Set(func(*sigmon.SignalMonitor) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil {
			log.Println(err)
		}

		fmt.Println("i'm melting!")
	})

	// listen and serve
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
	}

	fmt.Println("goodbye")
}
