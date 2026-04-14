// Handlers for menu-points
package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	svc "github.com/PhosFactum/KVStore/internal/service"
)

// I know, that this is bad for prod, but it's CLI-app (better use DI)
var globalStore = svc.NewStorage[string, string]()

// callSET: SET method handler
func CallSET(args []string) string {
	if len(args) < 2 {
		return "Usage: SET key value [TTL seconds]"
	}

	key := args[0]
	value := args[1]

	var ttl time.Duration

	if len(args) >= 4 && strings.ToUpper(args[2]) == "TTL" {
		seconds, err := strconv.Atoi(args[3])
		if err != nil {
			return fmt.Sprintf("Invalid TTL value: %s", args[3])
		}
		if seconds < 0 {
			return "TTL cannot be negative"
		}
		ttl = time.Duration(seconds) * time.Second
	}

	globalStore.SET(key, value, ttl)

	return "OK"
}

// callGET: GET method handler
func CallGET(args []string) string {
	if len(args) < 1 {
		return "Usage: GET key"
	}

	key := args[0]
	value, found := globalStore.GET(key)

	if !found {
		return "(nil)"
	}
	return fmt.Sprintf("'%s'", value)
}

// callDELETE: DELETE method handler
func CallDELETE(args []string) string {
	if len(args) < 1 {
		return "Usage: DELETE key"
	}

	key := args[0]
	wasDeleted := globalStore.DELETE(key)

	if wasDeleted {
		return "OK"
	}

	return "Key not found"
}

// CallSTATS: STATS method handler
func CallSTATS(args []string) string {
	stats := globalStore.STATS()

	return fmt.Sprintf(
		"Hits: %d, Misses: %d, Keys: %d, HitRate: %.2f%%",
		stats.Hits,
		stats.Misses,
		stats.Keys,
		stats.HitRate(),
	)
}

// GetStore: getter for main.go to storage from outer
func GetStore() *svc.Storage[string, string] {
	return globalStore
}
