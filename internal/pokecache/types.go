package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	c  map[string]cacheEntry
	mu *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	nc := Cache{
		c:  map[string]cacheEntry{},
		mu: &sync.Mutex{},
	}
	go nc.reapLoop(interval)
	return nc
}
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.c[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, ok := c.c[key]
	if ok {
		return data.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, entry := range c.c {
			if now.Sub(entry.createdAt) > interval {
				delete(c.c, key)
			}
		}
		c.mu.Unlock()
	}
}
