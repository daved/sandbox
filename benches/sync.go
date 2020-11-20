package main

import (
	"sync"
	"sync/atomic"
)

type safeBool struct {
	mu sync.Mutex
	v  bool
}

func (s *safeBool) value() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.v
}

func (s *safeBool) setValue(v bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.v = v
}

type atomicBool struct {
	n int32
}

func (a *atomicBool) value() bool {
	n := atomic.LoadInt32(&a.n)
	return n > 0
}

func (a *atomicBool) setValue(v bool) {
	var n int32
	if v {
		n = 1
	}
	atomic.StoreInt32(&a.n, n)
}
