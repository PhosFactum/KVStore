// Entrypoint of KVStore
package main

import "github.com/PhosFactum/KVStore/internal/app"

// main: entrypoint of program
func main() {
	app.NewApp().Run()
}
