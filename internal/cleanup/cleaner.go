// Cleaner realization
package cleanup

import (
	"sync"
	"time"
)

// CleanerInterface: contract to avoid cycle imports
type CleanerInterface interface {
	CleanupExpired() int
}

// Cleaner: managing background goroutine for clear
type Cleaner struct {
	store    CleanerInterface
	interval time.Duration
	stopCh   chan struct{} // chan for stop-signal
	wg       sync.WaitGroup
}

// NewCleaner: constructor for new cleaner creationg
func NewCleaner(interval time.Duration, store CleanerInterface) *Cleaner {
	return &Cleaner{
		store:    store,
		interval: interval,
		stopCh:   make(chan struct{}),
	}
}

// Start: starting background goroutine
func (c *Cleaner) Start() {
	c.wg.Add(1)
	go c.run()
}

// run: running goroutine
func (c *Cleaner) run() {
	defer c.wg.Done()

	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.store.CleanupExpired() // Ticker triggered - start cleanup
		case <-c.stopCh:
			return // stop-signal triggered - return and stop goroutine
		}
	}
}

// Stop: blocking completion until goroutine stop
func (c *Cleaner) Stop() {
	close(c.stopCh)
	c.wg.Wait()
}
