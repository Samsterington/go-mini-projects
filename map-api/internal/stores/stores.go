package stores

import (
	"fmt"
	"github.com/google/uuid"
	"go-mini-projects/map-api/internal/key_value_store"
	"sync"
)

type Stores struct {
	stores map[uuid.UUID]*key_value_store.KeyValueStore
	mutex  *sync.RWMutex
}

func Create() *Stores {
	return &Stores{
		stores: make(map[uuid.UUID]*key_value_store.KeyValueStore),
		mutex:  new(sync.RWMutex),
	}
}

func (s *Stores) GenerateNewStore() string {
	storeId := uuid.New()
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.stores[storeId] = key_value_store.Create()
	return storeId.String()
}

func (s *Stores) Insert(storeId string, key int, value string) (err error) {
	store, err := s.getStore(storeId)
	if err != nil {
		return err
	}

	store.Insert(key, value)
	return nil
}

func (s *Stores) Read(storeId string, key int) (value *string, err error) {
	store, err := s.getStore(storeId)
	if err != nil {
		return nil, err
	}

	value = store.Read(key)
	return value, nil
}

func (s *Stores) getStore(storeId string) (*key_value_store.KeyValueStore, error) {
	id, err := storeIdToUUID(storeId)
	if err != nil {
		return nil, fmt.Errorf("could not convert store id to uuid: %w", err)
	}

	s.mutex.RLock()
	store, ok := s.stores[id]
	s.mutex.RUnlock()

	if !ok {
		return nil, fmt.Errorf("store with id %s did not exist", id)
	}
	return store, nil
}

func storeIdToUUID(storeId string) (id uuid.UUID, err error) {
	id, err = uuid.Parse(storeId)
	if err != nil {
		return id, fmt.Errorf("store id is not valid uuid: %w", err)
	}
	return id, nil
}
