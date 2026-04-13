// Structure of storage and item
package store

import "time"

// KeyValueStorage: main storage structure
type KeyValueStorage[K comparable, V any] struct {
	storage map[K]*Item[V]
}

// Item: structure of each item in storage
type Item[V any] struct {
	Value      V
	Expiration time.Time
}
