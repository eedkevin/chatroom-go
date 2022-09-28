package inmemory

import (
	"fmt"
	"sync"
)

type InmemoryStorage struct {
	mutex sync.RWMutex
	data  map[string]interface{}
}

func NewStorage() *InmemoryStorage {
	return &InmemoryStorage{
		data: make(map[string]interface{}),
	}
}

func (s *InmemoryStorage) Save(id string, obj interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[id] = obj
	return nil
}

func (s *InmemoryStorage) Get(id string) (interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	val, ok := s.data[id]
	if !ok { // not exists
		return nil, nil
	}
	return val, nil
}

func (s *InmemoryStorage) List() (map[string]interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.data, nil
}

func (s *InmemoryStorage) Delete(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.data, id)
	return nil
}

func (s *InmemoryStorage) Update(id string, obj interface{}) (interface{}, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, ok := s.data[id]
	if !ok { // not exists
		return nil, fmt.Errorf("item[%s] not exists in storage", id)
	}
	s.data[id] = obj
	return obj, nil
}
