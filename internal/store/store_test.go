// Test of store.go
package store

import "testing"

// TestSET: testing of SET method
func TestSET(t *testing.T) {
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
