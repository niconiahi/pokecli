package main

import (
	"testing"
	"time"
)

type Test struct {
	key   string
	value []byte
}

func TestAddAndGetEntry(t *testing.T) {
	cache := createCache()

	tests := []Test{
		{
			key:   "key1",
			value: []byte("val1"),
		},
		{
			key:   "key1",
			value: []byte("val1"),
		},
	}

	for _, test := range tests {
		cache.Add(test.key, test.value)

		entry, ok := cache.Get(test.key)
		if !ok {
			t.Errorf("%s not found", test.key)
		}

		if string(entry) != string(test.value) {
			t.Errorf("values don't match")
		}
	}
}

func TestCreateCache(t *testing.T) {
	cache := createCache()

	if cache.entries == nil {
		t.Error("cache in nil")
	}
}

func TestPurgeSuccess(t *testing.T) {
	cache := createCache()
	duration := time.Millisecond * 10
	go cache.StartPurgeLoop(duration)

	key := "key1"
	value := "val1"
	cache.Add(key, []byte(value))

	time.Sleep(duration + time.Millisecond)

	_, ok := cache.Get(key)

	if ok {
		t.Errorf("%s should have been purged", key)
	}
}

func TestPurgeFail(t *testing.T) {
	cache := createCache()
	duration := time.Millisecond * 10
	go cache.StartPurgeLoop(duration)

	key := "key1"
	value := "val1"
	cache.Add(key, []byte(value))

	time.Sleep(duration / 2)

	_, ok := cache.Get(key)

	if !ok {
		t.Errorf("%s should not have been purged", key)
	}
}
