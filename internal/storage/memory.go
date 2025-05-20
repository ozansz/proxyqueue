package storage

import (
	"fmt"
	"sync"
	"time"
)

type memoryStorage[K comparable, V any] struct {
	store sync.Map
}

func NewMemoryStorage[K comparable, V any]() Storage[K, V] {
	return &memoryStorage[K, V]{
		store: sync.Map{},
	}
}

func (s *memoryStorage[K, V]) Set(key K, value V, _ time.Duration) error {
	s.store.Store(key, value)
	return nil
}

func (s *memoryStorage[K, V]) Get(key K) (V, error) {
	value, ok := s.store.Load(key)
	if !ok {
		var zero V
		return zero, fmt.Errorf("key not found")
	}
	return value.(V), nil
}

func (s *memoryStorage[K, V]) Close() error {
	return nil
}
