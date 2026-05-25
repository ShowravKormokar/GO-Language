## What is Redis

Redis (**REmote DIctionary Server**) is an open‑source, in‑memory key‑value data structure store that supports data types such as strings, lists, hashes, sets, and sorted sets. [web:4][web:7] It is commonly used as a **cache**, **message broker**, and lightweight **persistent datastore** for low‑latency operations. [web:1][web:7]

Redis is optimized for high‑performance reads and writes, typically serving data in microseconds from RAM while optionally persisting to disk for durability. [web:2][web:8]

---

## Why use Redis / Purpose

Redis is widely adopted because:

- **Speed**: data lives primarily in memory, enabling microsecond‑scale read/write operations. [web:1][web:7]  
- **Simplicity**: simple key‑based commands and a small set of well‑defined data types make it easy to learn and integrate. [web:6][web:7]  
- **Versatility**: use cases include:
  - Caching (e.g., HTTP responses, database query results).
  - Session stores for web applications.
  - Leaderboards (sorted sets).
  - Pub/Sub message‑oriented communication.
  - Rate‑limiting and job queues. [web:4][web:10]

---
<img width="716" height="1600" alt="WhatsApp Image 2026-05-26 at 12 27 27 AM" src="https://github.com/user-attachments/assets/8785e402-c681-46e7-aca5-77a684d3b048" />
---

## How Redis works internally (brief)

### In‑memory dataset

The primary dataset lives in **RAM**, backed by optional disk persistence mechanisms (RDB snapshots and AOF logging). [web:2][web:8] This design prioritizes **low latency** while allowing configuration for durability.

### Single‑threaded core

Redis executes commands in a **single‑threaded event loop**, which avoids complex locking and makes latency predictable. [web:2][web:10] Network I/O and background tasks (persistence, replication) are handled separately to keep the core simple.

### Optimized data structures

Each data type (string, list, hash, set, sorted set) is implemented in **highly optimized C code** that minimizes memory usage and CPU overhead. [web:3][web:6]

### Persistence and replication

- **RDB (Redis Database)**: periodic snapshots of the dataset written to disk. [web:5][web:8]  
- **AOF (Append‑Only File)**: every write command is appended to a log file, which can be replayed to rebuild the dataset. [web:2][web:5]  
- **Replication**: master‑replica replication supports high‑availability and read‑scaling setups. [web:8][web:10]

---

## Install Redis CLI / Server

### Windows (via WSL recommended)

Using **WSL + Ubuntu** is the most reliable way to run Redis on Windows.

```powershell
# In PowerShell (Windows)
wsl --install
wsl
```

```bash
# Inside WSL (Ubuntu)
sudo apt update
sudo apt install redis-server
sudo service redis-server start
redis-cli ping   # should reply PONG
```

This approach uses the native Linux Redis package via WSL, which is fully compatible. [web:7]

---

### macOS (Homebrew)

Install Redis with **Homebrew**:

```bash
brew update
brew install redis
brew services start redis
redis-cli ping   # PONG
```

Homebrew also supports `brew services stop redis` and `brew services restart redis` for control. [web:7]

---

### Linux (Ubuntu/Debian)

On Ubuntu/Debian‑based distributions:

```bash
sudo apt update
sudo apt install redis-server
sudo systemctl enable --now redis-server
redis-cli ping   # PONG
```

This installs the official Redis server package and enables it as a systemd service. [web:7]

---

## Basic Redis commands (quick reference)

| Command (example)              | Type     | Brief description |
|--------------------------------|----------|-------------------|
| `SET key value`               | String   | Set a key to a string value. [web:6][web:9] |
| `GET key`                     | String   | Get the value of a key. [web:6][web:9] |
| `DEL key`                     | Key      | Delete one or more keys. [web:6][web:9] |
| `RPUSH key val1 val2 ...`     | List     | Append values to the end of a list. [web:3][web:6] |
| `LPUSH key val1 val2 ...`     | List     | Prepend values to the start of a list. [web:3][web:6] |
| `LRANGE key 0 -1`             | List     | Read all elements of a list (from index 0 to last). [web:3][web:6] |
| `HSET key field value`        | Hash     | Set a field in a hash to a value. [web:3][web:6] |
| `HGET key field`              | Hash     | Get a single field from a hash. [web:3][web:6] |
| `HGETALL key`                 | Hash     | Get all fields and values of a hash. [web:3][web:6] |
| `SADD key member1 member2 ...`| Set      | Add one or more members to a set. [web:3][web:6] |
| `SMEMBERS key`                | Set      | Return all members of a set. [web:3][web:6] |
| `ZADD key score member`       | Sorted Set | Add a member with a score to a sorted set. [web:6][web:9] |
| `ZRANGE key 0 -1`             | Sorted Set | Get all members (with or without scores). [web:6][web:9] |
| `EXPIRE key seconds`          | Key      | Set a time‑to‑live (TTL) on a key. [web:6] |
| `PING`                        | Server   | Health check; returns `PONG` if server is alive. [web:7][web:9] |
| `KEYS *`                      | Key      | List all keys (use cautiously in production). [web:6] |

---

## Install Redis client for Go (step‑by‑step)

Redis provides an **official Go client** named `go-redis`, maintained in the `redis/go-redis` repository. [web:7][web:8]

### 1. Initialize a Go module

From your project root:

```bash
cd /path/to/your-project    # e.g., D:\GO-Language\Radis_Caching on Windows
go mod init radis_caching
go get github.com/redis/go-redis/v9
```

This creates `go.mod` and downloads the latest `v9` release of the official Go client. [web:7]

### 2. Minimal Go program connecting to Redis

Create `main.go`:

```go
package main

import (
  "context"
  "fmt"
  "github.com/redis/go-redis/v9"
)

func main() {
  ctx := context.Background()
  rdb := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
  })

  err := rdb.Set(ctx, "greeting", "hello from go", 0).Err()
  if err != nil {
    panic(err)
  }

  val, err := rdb.Get(ctx, "greeting").Result()
  if err != nil {
    panic(err)
  }

  fmt.Println(val)
}
```

Run:

```bash
go run .
```

This program connects to Redis running on `localhost:6379`, sets a key, retrieves it, and prints the value. [web:7]

### 3. Quick verification from redis-cli

From PowerShell or terminal:

```bash
redis-cli
```

Inside the CLI:

```redis
SET test "ok"
GET test
DEL test
```

If `GET test` returns `"ok"`, the Go client and Redis server are communicating correctly. [web:7]

---
## Screenshot:
<img width="753" height="293" alt="image" src="https://github.com/user-attachments/assets/abdf541c-6116-4117-b637-e35dabc39f70" />
---

## Next steps / Learning path

- Practice **Strings, Lists, Hashes, Sets, Sorted Sets** from Go using `go-redis`.  
- Learn about **TTL**, **eviction policies**, and **persistence (RDB/AOF)** for production‑ready setups. [web:5][web:8]  
- Explore **replication** and **clustering** for high‑availability and horizontal scaling. [web:8][web:10]

---

## References and official sources

- **Redis official documentation**: [https://redis.io/docs/latest/](https://redis.io/docs/latest/) [web:7][web:8]  
- **Redis GitHub (go-redis)**: [https://github.com/redis/go-redis](https://github.com/redis/go-redis) [web:7]  
- **Redis data types and quick start guide**: [https://redis.io/docs/latest/develop/get-started/data-store/](https://redis.io/docs/latest/develop/get-started/data-store/) [web:7]  
- **Redis persistence and durability (RDB & AOF)**: [https://redis.io/tutorials/operate/redis-at-scale/persistence-and-durability/](https://redis.io/tutorials/operate/redis-at-scale/persistence-and-durability/) [web:8]

---
## Made By [https://github.com/showravkormokar](Showrav Kormokar🤍)
