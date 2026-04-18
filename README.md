# KVStore
**Simple Key-Value storage like Redis as Pet-project written in Go**
 
## Features
- Thread-safe: 'sync.RWMutex' for data access, 'atomic.Int64' for metrics
- TTL support: automatic expiration with lazy-check + background cleaner
- Background cleanup: non-blocking worker with ticker, channels and waitgroups
- Runtime metrics: Hit/Miss counters and key count via 'STATS'
- Graceful Shutdown: proper goroutine termination on 'SIGTERM'
- Manual DI: composition root pattern in 'internal/app/'
- Race-free: all critical paths covered with '-race' tests (operations)


## Requirements
- Go version '1.26' or higher
- Make (optional, for convenience)


## Quick Start
``` bash
git clone https://github.com/PhosFactum/KVStore.git
cd KVStore
make run
```

Or run directly:
``` bash
go run cmd/app/main.go
```


## CLI Usage
1. Run program
2. Write in CLI commands like
``` bash
> SET key value
> GET key
> DELETE key
> STATS
> HELP
> EXIT
```


### Example Session
```
> SET session:abc token123 TTL 5
OK
> GET session:abc
'token123'
# Wait 6 seconds...
> GET session:abc
(nil)
> STATS
Hits: 1, Misses: 1, Keys: 0, HitRate: 50.00%
```

## Architecture
``` 
cmd/app/     - Entry point (1-liner)
internal/
|- app/      - Composition Root
|- cleanup/  - Background cleaner
|- handlers/ - CLI command routing
|- models/   - Domain structures
|- service/  - Core KVStore logic
|- ui/       - Interactive menu loop
pkg/         - Shared utils (input)
```


## Functions of KVStore
- Save data
- Get data
- Data have TTL (so u can save data with TTL)
- Delete data
- Get stats 
- Show helping window
