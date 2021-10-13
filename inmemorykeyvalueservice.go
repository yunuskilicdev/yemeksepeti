package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type InMemoryKV struct {
	data map[string]string
	sync.RWMutex
}

func (inMemory *InMemoryKV) Get(key string) string {
	inMemory.RLock()
	defer inMemory.RUnlock()
	return inMemory.data[key]
}

func (inMemory *InMemoryKV) Put(key string, value string) {
	inMemory.Lock()
	defer inMemory.Unlock()
	inMemory.data[key] = value
}

func (inMemory *InMemoryKV) DeleteAll() {
	inMemory.Lock()
	defer inMemory.Unlock()
	deleteAllFiles()
}

func (inMemory *InMemoryKV) persist() {
	inMemory.Lock()
	defer inMemory.Unlock()
	jsonStr, _ := json.Marshal(store.data)
	createFile(jsonStr)
}

func Store() *InMemoryKV {
	if store != nil {
		return store
	}
	store = &InMemoryKV{data: map[string]string{}}
	initializeBasePath()
	readJsonFromFile(&store.data)

	go store.backgroundTask()

	return store
}

func (inMemory *InMemoryKV) backgroundTask() {
	ticker := time.NewTicker(1 * time.Minute)
	for _ = range ticker.C {
		fmt.Println("PERSIST")
		store.persist()
	}
}

var store *InMemoryKV
