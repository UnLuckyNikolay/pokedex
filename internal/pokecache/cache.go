package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		cacheMap: map[string]cacheEntry{},
		mu:       &sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return &newCache
}

func (c *Cache) ListAll() {
	for key, item := range c.cacheMap {
		fmt.Printf(" > %v - %v\n", key, string(item.val))
	}
}

func (c *Cache) Add(urlKey string, rawData []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheMap[urlKey] = cacheEntry{
		createdAt: time.Now(),
		val:       rawData,
	}
}

func (c *Cache) Get(urlKey string) (rawData []byte, exists bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.cacheMap[urlKey]
	return v.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()

		currentTime := time.Now()
		for key, item := range c.cacheMap {
			if currentTime.Compare(item.createdAt.Add(interval)) == +1 {
				delete(c.cacheMap, key)
			}
		}

		c.mu.Unlock()
	}
}
