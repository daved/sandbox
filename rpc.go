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

func (r *RPC) Time(zone string, curTime *time.Time) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.reqs++

	loc, err := zones.LocByZone(zone)
	if err != nil {
		return ErrZoneNotFound
	}

	*curTime = time.Now().In(loc)

	return nil
}

func (r *RPC) Stats(_ bool, reqs *uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	*reqs = r.reqs

	return nil
}

func (r *RPC) Ping(_ bool, _ *bool) error {
	return nil
}
