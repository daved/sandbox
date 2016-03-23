package rpctime

import (
	"net"
	"net/rpc"
	"sync"
	"time"
)

type Client struct {
	mu      sync.Mutex
	conn    net.Conn
	rpcCli  *rpc.Client
	timeout time.Duration
}

func NewClient(dsn string, timeout time.Duration) (*Client, error) {
	conn, err := net.DialTimeout("tcp", dsn, timeout)
	if err != nil {
		return nil, err
	}

	c := &Client{
		conn:    conn,
		rpcCli:  rpc.NewClient(conn),
		timeout: timeout,
	}

	go func() {
		for {
			time.Sleep(time.Second * 6)

			c.manageConn()
		}
	}()

	return c, nil
}

func (c *Client) manageConn() {
	c.mu.Lock()
	defer c.mu.Unlock()

	var dummy bool
	if err := c.rpcCli.Call("RPC.Ping", false, &dummy); err != nil {
		c.conn.Close()

		conn, err := net.DialTimeout("tcp", c.conn.RemoteAddr().String(), c.timeout)
		if err != nil {
			return
		}

		c.conn = conn
		c.rpcCli = rpc.NewClient(conn)
	}
}

func (c *Client) Time(zone string) (time.Time, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var curTime *time.Time

	err := c.rpcCli.Call("RPC.Time", zone, &curTime)

	return *curTime, err
}

func (c *Client) Stats() (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var reqs uint64

	err := c.rpcCli.Call("RPC.Stats", false, &reqs)

	return reqs, err
}
