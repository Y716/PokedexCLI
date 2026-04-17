package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheEntry map[string]cacheEntry
	mu         *sync.Mutex
	duration   time.Duration
}

func NewCache(interval time.Duration) Cache {
	newCache := make(map[string]cacheEntry)
	cache := Cache{
		cacheEntry: newCache,
		duration:   interval,
		mu:         &sync.Mutex{},
	}
	fmt.Println("cache map initialized:", cache.cacheEntry != nil)
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newCacheEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.cacheEntry[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheEntry, ok := c.cacheEntry[key]
	if !ok {
		return nil, false
	}
	return cacheEntry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.duration)

	for range ticker.C {
		c.mu.Lock()

		for key, entry := range c.cacheEntry {
			if time.Since(entry.createdAt) > c.duration {
				delete(c.cacheEntry, key)
			}
		}
		c.mu.Unlock()
	}

}
