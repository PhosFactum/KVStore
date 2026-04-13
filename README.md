# **KVStore**
# **Simple Key-Value storage like Redis as Pet-project**

### **Architecture of project**

#### **internal/**: inner logic (hidden)
#### **store/**: main package 
- *store.go*: public API (Set, Get, Delete)
- *item.go*: element structure (value + TTL)
- *stats.go*: stats collecting
#### **tests/**: tests (unit+integration)
#### **pkg/**: optional utilities (errors)

#### **main.go**: entrypoing

### Functions of KVStore
- Save data
- Get data
- Data have TTL
- Delete data
- Mass operations (multiple operations)
- Monitoring

### How to use
1. Run program
2. Write in CLI commands like
``` bash
> SET key
> GET key
> STATS
> DELETE key
```
