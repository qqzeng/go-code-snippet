package maputil

import (
"sync"
)

type RawStringMap struct {
	mu sync.RWMutex
	internal map[string]string
}

func NewRawStringMap() *RawRWStringMap {
	return &RawRWStringMap{
		internal: make(map[string]string),
	}
}

func (rm *RawStringMap) Load(key string) (value string, ok bool) {
	rm.mu.RLock()
	result, ok := rm.internal[key]
	rm.mu.RUnlock()
	return result, ok
}

func (rm *RawStringMap) Delete(key string) {
	rm.mu.Lock()
	delete(rm.internal, key)
	rm.mu.Unlock()
}

func (rm *RawStringMap) Store(key, value string) {
	rm.mu.Lock()
	rm.internal[key] = value
	rm.mu.Unlock()
}

type RawIntMap struct {
	mu sync.RWMutex
	internal map[int]int
}

func NewRawIntMap() *RawRWIntMap {
	return &RawRWIntMap{
		internal: make(map[int]int),
	}
}

func (rm *RawIntMap) Load(key int) (value int, ok bool) {
	rm.mu.Lock()
	result, ok := rm.internal[key]
	rm.mu.Unlock()
	return result, ok
}

func (rm *RawIntMap) Delete(key int) {
	rm.mu.Lock()
	delete(rm.internal, key)
	rm.mu.Unlock()
}

func (rm *RawIntMap) Store(key, value int) {
	rm.mu.Lock()
	rm.internal[key] = value
	rm.mu.Unlock()
}
