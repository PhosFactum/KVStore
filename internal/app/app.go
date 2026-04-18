// DI-container
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PhosFactum/KVStore/internal/cleanup"
	"github.com/PhosFactum/KVStore/internal/handlers"
	"github.com/PhosFactum/KVStore/internal/service"
	"github.com/PhosFactum/KVStore/internal/ui"
)

// App: structure of main app
type App struct {
	store   *service.Storage[string, string]
	cleaner *cleanup.Cleaner
}

// NewApp: constructor for our app creation
func NewApp() *App {
	// Domain layer (core)
	store := service.NewStorage[string, string]()

	handlers.InitStore(store)

	cleaner := cleanup.NewCleaner(2*time.Second, store)

	return &App{
		store:   store,
		cleaner: cleaner,
	}
}

// Run: main app running goroutine
func (a *App) Run() {
	// Run background program
	a.cleaner.Start()

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// UI in separate goroutine
	go func() {
		ui.ShowMenu()
		sigChan <- syscall.SIGTERM
	}()

	<-sigChan

	fmt.Println("\nShutting down...")
	a.cleaner.Stop()
	fmt.Println("Stopped!")
}
