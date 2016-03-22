package rpctime

import (
	"errors"
	"sync"
	"time"

	"github.com/daved/rpctime/internal/zones"
)

var (
	ErrZoneNotFound = errors.New("zone not found")
)

type RPC struct {
	mu   sync.Mutex
	reqs uint64
}

func NewRPC() *RPC {
	return &RPC{}
}

func (r *RPC) Time(zone string, curTime *string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.reqs++

	loc, err := zones.LocByZone(zone)
	if err != nil {
		return ErrZoneNotFound
	}

	*curTime = time.Now().In(loc).String()

	return nil
}

func (r *RPC) Stats(skip bool, reqs *uint64) error {
	*reqs = r.reqs

	return nil
}
