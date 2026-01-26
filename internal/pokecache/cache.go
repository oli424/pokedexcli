package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	interval time.Duration
	mu       sync.Mutex
	items    map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		interval: interval,
		items:    make(map[string]cacheEntry),
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(url string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[url] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(url string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, ok := c.items[url]; ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for url, entry := range c.items {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.items, url)
			}
		}
		c.mu.Unlock()
	}
}
