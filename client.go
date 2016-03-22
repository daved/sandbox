package rpctime

import (
	"net"
	"net/rpc"
	"time"
)

type Client struct {
	conn *rpc.Client
}

func NewClient(dsn string, timeout time.Duration) (*Client, error) {
	conn, err := net.DialTimeout("tcp", dsn, timeout)
	if err != nil {

		return nil, err
	}

	c := &Client{
		conn: rpc.NewClient(conn),
	}

	return c, nil
}

func (c *Client) Time(zone string) (string, error) {
	var curTime string

	err := c.conn.Call("RPC.Time", zone, &curTime)

	return curTime, err
}

func (c *Client) Stats() (uint64, error) {
	var reqs uint64

	err := c.conn.Call("RPC.Stats", true, &reqs)

	return reqs, err
}
