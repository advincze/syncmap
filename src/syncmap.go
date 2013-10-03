package syncmap

import (
	"sync"
)

type SyncedMap struct {
	mutex sync.RWMutex
	m     map[interface{}]interface{}
	in    chan struct{ k, v interface{} }
}

func NewSyncedMap() *SyncedMap {
	syncedMap := &SyncedMap{m: make(map[interface{}]interface{}), in: make(chan struct{ k, v interface{} })}
	go syncedMap.writeLoop()
	return syncedMap
}

func (s *SyncedMap) Set(key interface{}, value interface{}) {
	s.in <- struct{ k, v interface{} }{key, value}
}

func (s *SyncedMap) writeLoop() {
	for {
		select {
		case kv := <-s.in:
			s.mutex.Lock()
			s.m[kv.k] = kv.v
			s.mutex.Unlock()
		}
	}
}

func (s *SyncedMap) Get(key interface{}) (interface{}, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	val, ok := s.m[key]
	return val, ok
}

func (s *SyncedMap) Stop() {
	close(s.in)
}
