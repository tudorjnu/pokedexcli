package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	interval time.Duration
	mu       sync.RWMutex
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
	c.mu.Lock()
	c.cacheMap[key] = cacheEntry{
		createdAt: now,
		val:       val,
	}
	c.mu.Unlock()
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.cacheMap[key]
	if !ok {
		return nil, false
	}
	v.createdAt = time.Now()

	return v.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		t, ok := <-ticker.C
		if ok {
			c.mu.Lock()
			for k, v := range c.cacheMap {
				if t.Sub(v.createdAt) > interval {
					delete(c.cacheMap, k)
				}
			}
			c.mu.Unlock()
		}
	}

}
