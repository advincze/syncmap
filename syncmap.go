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

func (s *SyncedMap) Remove(key interface{}) (interface{}, bool){
  s.Lock()
  defer s.Unlock()
  if s.m[key] != nil {
    delete(s.m, key)
    return key, true
  }
  return key, false
}

func (s *SyncedMap) Iterate(f func(interface{}, interface{})) {
  for key := range s.m {
    val, _ := s.Get(key)
    f(key, val)
  }
}
