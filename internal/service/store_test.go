// Test of store.go
package store

import (
	"sync"
	"testing"
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
