package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/codemodus/renode"
	"github.com/codemodus/sigmon"
)

type server struct {
	*renode.Node
	dataMu sync.Mutex
	data   int
}

func newServer(port string) *server {
	s := &server{}

	m := http.NewServeMux()
	m.HandleFunc("/", s.handler)

	opts := renode.Options{
		Timeout: time.Second,
		Addr:    port,
		Handler: m,
	}
	n, err := renode.New(opts)
	if err != nil {
		fmt.Println("YIKES - Return an error instead!")
	}

	s.Node = n
	return s
}

func (s *server) handler(w http.ResponseWriter, r *http.Request) {
	s.dataMu.Lock()
	defer s.dataMu.Unlock()

	fmt.Fprintf(w, "restart count: %d\n", s.data)
}

func (s *server) restart() {
	s.dataMu.Lock()
	defer s.dataMu.Unlock()

	s.data++
}

func (s *server) signalHandler(sm *sigmon.SignalMonitor) {
	switch sm.Sig() {
	case sigmon.SIGINT, sigmon.SIGTERM:
		t := time.Now()
		s.Stop()
		fmt.Printf("%-32s [%s] %v\n", "handled signal", sm.Sig(), time.Since(t))

	case sigmon.SIGHUP:
		t := time.Now()
		s.Restart(s.restart)
		fmt.Printf("%-32s [%s] %v\n", "handled signal", sm.Sig(), time.Since(t))

	case sigmon.SIGUSR1, sigmon.SIGUSR2:
		fmt.Println("ignored")
	}
}

func main() {
	sm := sigmon.New(nil)
	sm.Run() // signals will not affect app

	s := newServer(":6464")
	s.Start()
	sm.Set(s.signalHandler) // signals will be handled

	err := s.Wait()
	if err != nil {
		fmt.Println(err)
	}

	sm.Stop()
}
