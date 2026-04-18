// Test of store.go
package service

import (
	"sync"
	"testing"
	"time"
)

// TestSETandGET: test for SET and GET methods
func TestSETandGET(t *testing.T) {
	// Arrange - creating testdata
	store := NewStorage[string, int]()

	// Act - writing data
	store.SET("user:1", 20, 0)

	// Assert - reading and checking
	value, found := store.GET("user:1")

	if !found {
		t.Fatalf("The key 'user:1' wasn't found!")
	}
	if value != 20 {
		t.Errorf("Expected 20, got %d", value)
	}
}

// TestDELETE: test for DELETE method
func TestDELETE(t *testing.T) {
	store := NewStorage[string, int]()

	store.SET("key-to-delete", 123, 0)

	wasDeleted := store.DELETE("key-to-delete")

	if !wasDeleted {
		t.Error("Expected TRUE, because key existed")
	}

	_, found := store.GET("key-to-delete")
	if found {
		t.Error("Expected FALSE, because key should was deleted")
	}

	wasDeletedAgain := store.DELETE("non-existent-key")
	if wasDeletedAgain {
		t.Error("Expected FALSE, because key wasn't exist")
	}
}

// TestSTATS: test of metrics collection
func TestSTATS(t *testing.T) {
	store := NewStorage[string, string]()

	// Empty at start
	stats := store.STATS()
	if stats.Hits != 0 || stats.Misses != 0 {
		t.Errorf("Expected empty stats, got %+v", stats)
	}

	// Hits
	store.SET("k1", "v1", 0)
	store.GET("k1")
	store.SET("k2", "k2", 0)
	store.GET("k2")

	// Miss (no key)
	store.GET("no-such-key")

	// Miss (TTL expired)
	store.SET("temp", "val", 1*time.Millisecond)
	time.Sleep(10 * time.Millisecond)
	store.GET("temp")

	stats = store.STATS()
	if stats.Hits != 2 {
		t.Errorf("Expected 2 hit, got %d", stats.Hits)
	}
	if stats.Misses != 2 {
		t.Errorf("Expected 2 misses, got %d", stats.Misses)
	}
	if rate := stats.HitRate(); rate != 50 {
		t.Logf("HitRate: %.2f%% (ok)", rate)
	}
}

// / --- Concurrent Tests --- ///
//
// TestConcurrentSET: test for SET and GET concurrently
func TestConcurrentSET(t *testing.T) {
	store := NewStorage[string, int]()
	wg := sync.WaitGroup{}

	// Writing
	for i := range 100 {
		wg.Go(func() {
			key := "key"
			store.SET(key, i, 0)
		})
	}

	// Reading
	for range 100 {
		wg.Go(func() {
			store.GET("key")
		})
	}

	wg.Wait()
}
