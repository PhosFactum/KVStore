// Public API (for Set, Get and Delete operations)
package service

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/PhosFactum/KVStore/internal/cleanup"
	m "github.com/PhosFactum/KVStore/internal/models"
)

// Storage: main storage structure
type Storage[K comparable, V any] struct {
	mtx  sync.RWMutex
	data map[K]*m.Item[V]

	// Statisтолько [одна фича]. Напиши структуру файлов и минимальный код для неё. Без объяснений. tics
	hits   atomic.Int64
	misses atomic.Int64

	// Background cleaner
	cleaner *cleanup.Cleaner
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
		s.misses.Add(1)
		var zero V
		return zero, false
	}

	// Check if item expired
	if !item.Expiration.IsZero() && time.Now().After(item.Expiration) {
		s.misses.Add(1)
		var zero V
		return zero, false
	}

	s.hits.Add(1)
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

// STATS: method for get STATS of storage
func (s *Storage[K, V]) STATS() m.Stats {
	s.mtx.RLock()
	keyCount := len(s.data)
	s.mtx.RUnlock()

	return m.Stats{
		Hits:   s.hits.Load(),
		Misses: s.misses.Load(),
		Keys:   keyCount,
	}
}

// CleanupExpired: deleting expired elements
func (s *Storage[K, V]) CleanupExpired() int {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	count := 0
	now := time.Now()
	for k, item := range s.data {
		if !item.Expiration.IsZero() && now.After(item.Expiration) { // TTL checking
			delete(s.data, k)
			count++
		}
	}

	return count
}

// StartCleaner: running automatic cleanup with interval
func (s *Storage[K, V]) StartCleaner(interval time.Duration) {
	s.cleaner = cleanup.NewCleaner(interval, s)
	s.cleaner.Start()
}

// StopCleaner: gracefully shutdown the goroutine
func (s *Storage[K, V]) StopCleaner() {
	if s.cleaner != nil {
		s.cleaner.Stop()
	}
}
