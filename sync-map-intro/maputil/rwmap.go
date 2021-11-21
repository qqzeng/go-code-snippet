package maputil

import (
"sync"
)

type RawRWStringMap struct {
	mu sync.RWMutex
	internal map[string]string
}

func NewRawRWStringMap() *RawRWStringMap {
	return &RawRWStringMap{
		internal: make(map[string]string),
	}
}

func (rm *RawRWStringMap) Load(key string) (value string, ok bool) {
	rm.mu.RLock()
	result, ok := rm.internal[key]
	rm.mu.RUnlock()
	return result, ok
}

func (rm *RawRWStringMap) Delete(key string) {
	rm.mu.Lock()
	delete(rm.internal, key)
	rm.mu.Unlock()
}

func (rm *RawRWStringMap) Store(key, value string) {
	rm.mu.Lock()
	rm.internal[key] = value
	rm.mu.Unlock()
}

type RawRWIntMap struct {
	mu sync.RWMutex
	internal map[int]int
}

func NewRawRWIntMap() *RawRWIntMap {
	return &RawRWIntMap{
		internal: make(map[int]int),
	}
}

func (rm *RawRWIntMap) Load(key int) (value int, ok bool) {
	rm.mu.RLock()
	result, ok := rm.internal[key]
	rm.mu.RUnlock()
	return result, ok
}

func (rm *RawRWIntMap) Delete(key int) {
	rm.mu.Lock()
	delete(rm.internal, key)
	rm.mu.Unlock()
}

func (rm *RawRWIntMap) Store(key, value int) {
	rm.mu.Lock()
	rm.internal[key] = value
	rm.mu.Unlock()
}
