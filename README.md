# rpctime

    go get github.com/daved/rpctime...

Package rpctime is an RPC client and server library, and also includes an 
implementation of the server within `./cmd/rpctime`. It is a trivial example 
and does not resolve many common concerns with production-ready microservices.

## Usage (Client/Server Library)

```go
type Client
    func NewClient(dsn string, timeout time.Duration) (*Client, error)
    func (c *Client) Stats() (uint64, error)
    func (c *Client) Time(zone string) (string, error)
type RPC
    func NewRPC() *RPC
    func (r *RPC) Ping(_ bool, _ *bool) error
    func (r *RPC) Stats(_ bool, reqs *uint64) error
    func (r *RPC) Time(zone string, curTime *string) error
```

## Usage (Server)

    Available flags (defaults in quotes):

    --port={:port} // ":19876"

## More Info

### Additional Code

An example implementation of the client code is available at 
https://github.com/daved/rpctimehttp.
