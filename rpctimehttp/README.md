# rpctimehttp

    go get github.com/daved/rpctimehttp

rpctimehttp is an HTTP API server making use of the 
https://github.com/daved/rpctime RPC server and client library.  It is a 
trivial example and does not resolve many common concerns with production-ready 
apps or microservices.

## Usage

    Available flags (defaults in quotes):
    
    --http-port={:port} // ":29876"
    --rpc-port={:port} // ":19876"
