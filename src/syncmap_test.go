package syncmap

import (
	"math/rand"
	"runtime"
	"testing"
)

func TestFirst(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	syncedMap := NewSyncedMap()

	syncedMap.Set(1, 1)
	val, ok := syncedMap.Get(1)

	if !ok {
		t.Errorf("value could not be retrieved\n")
	}
	if val != 1 {
		t.Errorf("value expected:%v but was %v\n", 1, val)
	}

	SetTest(syncedMap)
	syncedMap.Stop()
}

func SetTest(syncedMap *SyncedMap) {
	for i := 0; i < 100; i++ {
		r := rand.Intn(100)
		syncedMap.Set(r, r)
	}
}
