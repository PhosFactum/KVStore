// Structure of item
package models

import "time"

// Item: structure of each item in storage
type Item[V any] struct {
	Value      V
	Expiration time.Time
}

// NewItem: constructor for new item
func NewItem[V any](value V, ttl time.Duration) *Item[V] {
	item := &Item[V]{Value: value}

	if ttl > 0 {
		item.Expiration = time.Now().Add(ttl)
	}

	return item
}
