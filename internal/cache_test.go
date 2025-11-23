package internal

import (
	"testing"
	"time"
)

func TestCacheAddAndGet(t *testing.T) {
	ttl := 1
	cache := InitCache(ttl)

	key := "location-1"
	value := []byte("some-bytes")

	added := cache.Add(key, value)
	if !added {
		t.Fatalf("expected Add to return true for new entry")
	}

	exists, entry := cache.Get(key)
	if !exists {
		t.Fatalf("expected entry to exist after Add")
	}

	if string(entry.value) != string(value) {
		t.Fatalf("expected value %q, got %q", string(value), string(entry.value))
	}
}

func TestCacheExpiration(t *testing.T) {
	ttl := 1
	cache := InitCache(ttl)

	key := "location-2"
	value := []byte("temp-value")

	added := cache.Add(key, value)
	if !added {
		t.Fatalf("expected Add to return true for new entry")
	}

	// Wait long enough for the entry to expire.
	time.Sleep(time.Duration(ttl * 2) * time.Second)

	exists, _ := cache.Get(key)
	if exists {
		t.Fatalf("expected entry to be expired and removed from cache")
	}
}
