package pokecache
import (
	"time"
	"sync"
)
type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	interval time.Duration
	mux *sync.RWMutex
	entries map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache{

	mux := &sync.RWMutex{}
	entry := map[string]cacheEntry{}
	cache := Cache{interval:interval, mux: mux, entries:entry }
	//go cache.Add(key, val, mux)
	go cache.readLoop()
	return cache
}

func (c Cache) Add(key string, val []byte) bool {
	c.mux.Lock()
	c.entries[key] = cacheEntry{createdAt:time.Now(), val:val}
	_, ok :=c.entries[key]
	c.mux.Unlock()
	return ok
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	entry, ok :=c.entries[key]
	c.mux.RUnlock()
	return entry.val, ok 
}

func (c Cache) readLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for _ = range ticker.C {
		for key, _ := range c.entries {
				if time.Since(c.entries[key].createdAt) > c.interval {
					c.mux.Lock()
					delete(c.entries, key)
					c.mux.Unlock()
				}
			}
	}
	
}
