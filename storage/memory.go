package storage

import (
	"encoding/hex"
	"errors"
	fmt "fmt"
	"sync"
)

type MemoryAdapter struct {
	store map[string][]byte
	lock  *sync.RWMutex
}

func NewMemoryAdapter() *MemoryAdapter {
	return &MemoryAdapter{
		store: make(map[string][]byte),
		lock:  &sync.RWMutex{},
	}
}

func (kv *MemoryAdapter) Get(key []byte) ([]byte, error) {
	kv.lock.RLock()
	defer kv.lock.RUnlock()
	keyHex := hex.EncodeToString(key)
	if v, ok := kv.store[keyHex]; ok {
		return v, nil
	} else {
		return nil, errors.New(fmt.Sprintf("[MemKV] key not found: %s", keyHex))
	}
}

func (kv *MemoryAdapter) Put(key, value []byte) error {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	keyHex := hex.EncodeToString(key)
	kv.store[keyHex] = value
	return nil
}

func (kv *MemoryAdapter) Has(key []byte) bool {
	kv.lock.RLock()
	defer kv.lock.RUnlock()
	keyHex := hex.EncodeToString(key)
	_, ok := kv.store[keyHex]
	return ok
}

func (kv *MemoryAdapter) Delete(key []byte) error {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	keyHex := hex.EncodeToString(key)
	if _, ok := kv.store[keyHex]; ok {
		delete(kv.store, keyHex)
	} else {
		return errors.New(fmt.Sprintf("[MemKV] key not found: %s", keyHex))
	}
	return nil
}

func (kv *MemoryAdapter) BatchPut(kvs [][2][]byte) error {
	kv.lock.Lock()
	defer kv.lock.Unlock()
	for i := range kvs {
		kv.Put(kvs[i][0], kvs[i][1])
	}
	return nil
}

func (kv *MemoryAdapter) Close() {}
