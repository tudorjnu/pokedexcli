package pokecache

import (
	"sync"
	"time"
)

// TODO: add mutex
type Cache struct {
	cacheMap map[string]cacheEntry
	interval time.Duration
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	m := make(map[string]cacheEntry)
	cache := Cache{
		cacheMap: m,
		interval: interval,
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	now := time.Now()
	c.cacheMap[key] = cacheEntry{
		createdAt: now,
		val:       val,
	}
}
func (c *Cache) Get(key string) ([]byte, bool) {
	v, ok := c.cacheMap[key]
	if !ok {
		return nil, false
	}

	return v.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		t, ok := <-ticker.C
		if ok {
			for k, v := range c.cacheMap {
				if t.Sub(v.createdAt) > interval {
					delete(c.cacheMap, k)
				}
			}
		}
	}

}
