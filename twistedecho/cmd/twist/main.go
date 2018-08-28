package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

type opts struct {
	port string
}

type twister struct {
	*net.TCPListener
}

func newTwister(port string) (*twister, error) {
	a, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		return nil, err
	}

	l, err := net.ListenTCP("tcp", a)
	if err != nil {
		return nil, err
	}

	return &twister{TCPListener: l}, nil
}

func (t *twister) run() error {
	for {
		t.connect()
	}
}

func (t *twister) connect() {
	conn, err := t.AcceptTCP()
	if err != nil {
		//return err
	}

	go t.process(conn)
}

func (t *twister) process(conn *net.TCPConn) {
	defer func() {
		conn.Close()
	}()

	for {
		//s := bufio.NewScanner(t)
		r := bufio.NewReader(conn)
		b, err := r.ReadBytes('\n')
		if err != nil {
			return //return err
		}

		if _, err := conn.Write(b); err != nil {
			return //return err
		}
	}
}

func main() {
	o := opts{}
	{
		flag.StringVar(&o.port, "port", ":31222", "port to listen on")
	}
	flag.Parse()

	t, err := newTwister(o.port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot get new twister: %v\n", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		if err := t.run(); err != nil {
			fmt.Fprintf(os.Stderr, "error running twister: %v\n", err)
		}
		wg.Done()
	}()

	wg.Wait()
}
