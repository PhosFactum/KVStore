// Public API (for Set, Get and Delete operations)
package store

import (
	"sync"
	"time"

	m "github.com/PhosFactum/KVStore/internal/models"
)

// Storage: main storage structure
type Storage[K comparable, V any] struct {
	mtx  sync.RWMutex
	data map[K]*m.Item[V]
}

// NewStorage: constructor for new storage creation
func NewStorage[K comparable, V any]() *Storage[K, V] {
	return &Storage[K, V]{
		data: make(map[K]*m.Item[V]),
	}
}

// SET: method for SET a value by key
func (s *Storage[K, V]) SET(key K, value V, ttl time.Duration) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.data[key] = m.NewItem(value, ttl)
}

// GET: method for GET a value by key
func (s *Storage[K, V]) GET(key K) (V, bool) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	item, ok := s.data[key]
	if !ok {
		var zero V
		return zero, false
	}

	// Check if item expired
	if !item.Expiration.IsZero() && time.Now().After(item.Expiration) {
		var zero V
		return zero, false
	}

	return item.Value, true
}

// DELETE: method for DELETE a value by key
func (s *Storage[K, V]) DELETE(key K) bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	_, exists := s.data[key]
	if exists {
		delete(s.data, key)
	}

	return exists
}
