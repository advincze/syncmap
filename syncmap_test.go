package syncmap

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

func TestFirst(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	syncedMap := NewSyncedMap()

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			setAndGetTest(t, syncedMap, 100)
      iterateTest(t, syncedMap)
      removeTest(t, syncedMap)
			wg.Done()
		}()
	}

	wg.Wait()
}

func setAndGetTest(t *testing.T, syncedMap *SyncedMap, n int) {
	for i := 0; i < n; i++ {
		r := rand.Intn(100)
		syncedMap.Set(r, r)
		val, ok := syncedMap.Get(r)
		if !ok {
			t.Errorf("value could not be retrieved\n")
		}
		if val != r {
			t.Errorf("value expected:%v but was %v\n", 1, val)
		}
	}
}

func removeTest(t *testing.T, syncedMap *SyncedMap) {
  r := rand.Uint32()
  syncedMap.Set(r,r)
  val, ok := syncedMap.Get(r)
  if !ok {
    t.Errorf("value could not be retrieved\n")
  }
  val, ok = syncedMap.Remove(r)
  if !ok {
    t.Errorf("value could not be removed\n")
  }
}

func iterateTest(t *testing.T, syncedMap *SyncedMap) {
  count := 0
  syncedMap.Iterate(func(key, value interface{}) {
    if key != nil {
      count += 1
    }
  })
  if count == 0 {
    t.Errorf("count is %d", count)
  }
}
