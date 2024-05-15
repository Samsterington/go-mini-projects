package key_value_store

import (
	"errors"
	"fmt"
	"sync"
)

type KeyValueStore struct {
	store map[int]string
	mutex *sync.RWMutex
}

func Create() *KeyValueStore {
	return &KeyValueStore{
		store: make(map[int]string),
		mutex: new(sync.RWMutex),
	}
}

func (k *KeyValueStore) Insert(key int, value string) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	k.store[key] = value
}

func (k *KeyValueStore) Read(key int) (string, bool) {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	value, exists := k.store[key]
	return value, exists
}

func (k *KeyValueStore) ReadRange(start int, end int) (map[int]string, error) {
	err := validateParams(start, end)
	if err != nil {
		return nil, fmt.Errorf("validating params: %w", err)
	}

	rangeLength := end - start
	if len(k.store) < rangeLength {
		return k.store, nil
	}

	subStore := make(map[int]string)
	for i := 0; i < rangeLength; i++ {
		value, exists := k.Read(i)
		if !exists {
			continue
		}
		subStore[i] = value
	}
	return subStore, nil
}

func validateParams(start int, end int) (err error) {
	if start > end {
		return errors.New("start must be before end")
	}
	return err
}
