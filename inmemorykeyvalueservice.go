package main

import "sync"

type InMemoryKV struct {
	data map[string]string
	mu   sync.Mutex
}

func (inMemory *InMemoryKV) Get(key string) string {
	inMemory.mu.Lock()
	defer inMemory.mu.Unlock()
	return inMemory.data[key]
}

func (inMemory *InMemoryKV) Put(key string, value string) {
	inMemory.mu.Lock()
	defer inMemory.mu.Unlock()
	inMemory.data[key] = value
}

func Store() *InMemoryKV {
	if store != nil {
		return store
	}
	store = &InMemoryKV{data: map[string]string{}}
	return store
}

var store *InMemoryKV
