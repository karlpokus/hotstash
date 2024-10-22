package main

import (
	"bytes"
	"sync"
)

// Stash is a thread-safe key-value store
type Stash struct {
	sync.RWMutex
	// key-value store protected by the mutex
	store map[string][]byte
}

// Hit returns true if k is mapped to v. If k exist but does not
// map to v, or if k is missing, k is mapped to v.
func (s *Stash) Hit(k string, v []byte) bool {
	if s.hit(k, v) {
		return true
	}
	s.push(k, v)
	return false
}

// hit returns true if k is mapped to v
func (s *Stash) hit(k string, v []byte) bool {
	s.Lock()
	defer s.Unlock()
	// key missing
	if _, ok := s.store[k]; !ok {
		return false
	}
	// value mismatch
	if !bytes.Equal(s.store[k], v) {
		return false
	}
	return true
}

// push pushes k and v to the store
func (s *Stash) push(k string, v []byte) {
	s.Lock()
	defer s.Unlock()
	s.store[k] = v
}

// newStash returns a new Stash
func newStash() *Stash {
	return &Stash{
		store: make(map[string][]byte),
	}
}
