package cache

import (
	"errors"
	"sync"
	"time"
)

type CacheItem struct {
	Value     string
	ExpiresAt time.Time // zero value , no expiration.
}

type CacheMetrics struct {
	Hits      int64
	Misses    int64
	ItemCount int64
}

type Cache struct {
	items  map[string]CacheItem
	mu     sync.RWMutex
	hits   int64
	misses int64
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]CacheItem),
	}
}

func (c *Cache) Set(key, value string, ttl int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	expiration := time.Time{}
	if ttl > 0 {
		expiration = time.Now().Add(time.Duration(ttl) * time.Second)
	}
	c.items[key] = CacheItem{
		Value:     value,
		ExpiresAt: expiration,
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	item, found := c.items[key]
	c.mu.RUnlock()

	if !found {
		c.mu.Lock()
		c.misses++
		c.mu.Unlock()
		return "", false
	}

	if !item.ExpiresAt.IsZero() && time.Now().After(item.ExpiresAt) {
		c.mu.Lock()
		delete(c.items, key)
		c.misses++
		c.mu.Unlock()
		return "", false
	}

	c.mu.Lock()
	c.hits++
	c.mu.Unlock()
	return item.Value, true
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, exists := c.items[key]; !exists {
		return errors.New("key not found")
	}
	delete(c.items, key)
	return nil
}

func (c *Cache) GetMetrics() CacheMetrics {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return CacheMetrics{
		Hits:      c.hits,
		Misses:    c.misses,
		ItemCount: int64(len(c.items)),
	}
}

func (c *Cache) CleanupExpired() int {
	removed := 0
	now := time.Now()
	c.mu.Lock()
	for key, item := range c.items {
		if !item.ExpiresAt.IsZero() && now.After(item.ExpiresAt) {
			delete(c.items, key)
			removed++
		}
	}
	c.mu.Unlock()
	return removed
}
