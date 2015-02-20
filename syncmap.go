package syncmap

import (
	"sync"
)

type SyncedMap struct {
	sync.RWMutex
	m map[interface{}]interface{}
}

func NewSyncedMap() *SyncedMap {
	return &SyncedMap{m: make(map[interface{}]interface{})}
}

func (s *SyncedMap) Set(key interface{}, value interface{}) {
	s.Lock()
	defer s.Unlock()

	s.m[key] = value
}

func (s *SyncedMap) Get(key interface{}) (interface{}, bool) {
	s.RLock()
	defer s.RUnlock()

	val, ok := s.m[key]
	return val, ok
}

func (s *SyncedMap) Iterate(f func(interface{}, interface{})) {
  for key := range s.m {
    f(key, s.m[key])
  }
}
