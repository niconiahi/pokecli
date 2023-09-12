package main

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]Entry
	mux     *sync.Mutex
}

type Entry struct {
	val       []byte
	createdAt time.Time
}

func createCache() Cache {
	cache := Cache{
		entries: make(map[string]Entry),
		mux:     &sync.Mutex{},
	}

	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	cache.entries[key] = Entry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	entry, ok := cache.entries[key]
	return entry.val, ok
}

func (cache *Cache) StartPurgeLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	for range ticker.C {
		cache.Purge(duration)
	}
}

func (cache *Cache) Purge(duration time.Duration) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	time := time.Now().UTC().Add(-duration)
	for i, entry := range cache.entries {
		if entry.createdAt.Before(time) {
			delete(cache.entries, i)
		}
	}
}
