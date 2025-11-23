package internal

import (
	"fmt"
	"time"
)

type cacheEntry struct {
	createdAt 	time.Time
	value 		[]byte
}

type Pokecache interface {
	Add(name string, val []byte) bool
	Get(name string) cacheEntry
	isExpired(name string) bool
	cleanup()
}

type Cache struct {
	entries map[string]cacheEntry
	ttl 	int
}

func InitCache(ttl int) *Cache {
	return &Cache{
		entries: make(map[string]cacheEntry),
		ttl: ttl,
	}
}

func(c Cache) isExpired(name string) bool {
	expiresAt := c.entries[name].createdAt.Add(time.Duration(c.ttl) * time.Second)

	return time.Now().After(expiresAt)
}

func(c *Cache) Add(name string, val []byte) bool {
	defer c.cleanup()

	if _, ok := c.entries[name]; ok {
		if c.isExpired(name) {
			return false
		}
	}

	fmt.Println("Adding to cache:")
	c.entries[name] = cacheEntry{
		value: val,
		createdAt: time.Now(),
	}

	return true
}

func(c Cache) Get(name string) (exists bool, entry cacheEntry) {
	c.cleanup()

	if entry, ok := c.entries[name]; ok {
		if c.isExpired(name) {
			delete(c.entries, name)
			return false, cacheEntry{}
		} else {
			fmt.Println("Getting from cache:")
			return true, entry
		}
	} else {
		return false, cacheEntry{}
	}
}

func(c *Cache) cleanup() {
	for key := range c.entries {
		if c.isExpired(key) {
			delete(c.entries, key)
		}
	}
}