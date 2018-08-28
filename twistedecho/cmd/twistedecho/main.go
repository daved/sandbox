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

type twistedConn struct {
	net.Conn
}

func newTwistedConn(port string) (*twistedConn, error) {
	conn, err := net.Dial("tcp", port)
	if err != nil {
		return nil, err
	}

	return &twistedConn{Conn: conn}, nil
}

func (c *twistedConn) run() error {
	for {
		r := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")

		b, err := r.ReadBytes('\n')
		if err != nil {
			return err
		}

		if _, err = c.Write(b); err != nil {
			return err
		}

		rr := bufio.NewReader(c)
		rb, err := rr.ReadBytes('\n')
		if err != nil {
			return err
		}

		_, err = os.Stdout.Write(rb)
		if err != nil {
			return err
		}
	}
}

func (c *twistedConn) stop() error {
	return c.Close()
}

func main() {
	o := opts{}
	{
		flag.StringVar(&o.port, "port", ":31222", "port of twisted process")
	}
	flag.Parse()

	c, err := newTwistedConn(o.port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot get new connection to twisted: %v\n", err)
		os.Exit(1)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		if err := c.run(); err != nil {
			fmt.Fprintf(os.Stderr, "error running connection to twisted: %v\n", err)
			os.Exit(1)
		}
		wg.Done()
	}()

	wg.Wait()
}
