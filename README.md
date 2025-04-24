# GoRedis

A lightweight Redis server implementation in Go that implements the Redis Serialization Protocol (RESP).

## Features

Currently supported Redis commands:

- `PING` - Test server connection
- `SET key value` - Set a key-value pair
- `GET key` - Get the value of a key
- `HSET hash field value` - Set a field in a hash
- `HGET hash field` - Get a field from a hash
- `HGETALL hash` - Get all fields and values from a hash

## Getting Started

### Prerequisites

- Go 1.23.5 or later

### Running the Server

```bash
go run . --port 6379  # Default port is 6379 if not specified
```

You can then connect to the server using any Redis client, including the official `redis-cli`.

## Implementation Details

- Uses Go's standard library for TCP networking
- Implements basic RESP (Redis Serialization Protocol) parsing and encoding
- Thread-safe operations using sync.RWMutex
- Supports both simple strings and bulk strings in the protocol
- Handles array types for command processing

## Potential Improvements

1. Data Persistence
   - [x] Add support for saving data to disk
   - [x] Implement Redis RDB-style persistence
   - [x] Add AOF (Append Only File) logging

2. Additional Commands
   - [ ] Add support for Lists (LPUSH, RPUSH, LRANGE, etc.)
   - [ ] Add support for Sets (SADD, SMEMBERS, etc.)
   - [ ] Add support for Sorted Sets (ZADD, ZRANGE, etc.)
   - [ ] Implement expiration for keys (EXPIRE, TTL)

3. Enhanced Features
   - [ ] Add support for pub/sub messaging
   - [ ] Implement transactions (MULTI, EXEC, WATCH)
   - [ ] Add authentication (AUTH command)
   - [ ] Add database selection (SELECT command)

4. Performance Improvements
   - [ ] Implement connection pooling
   - [ ] Add support for pipelining
   - [ ] Optimize memory usage
   - [ ] Add benchmarking tools

5. Operational Improvements
   - [ ] Add proper logging and monitoring
   - [ ] Implement graceful shutdown
   - [ ] Add configuration file support
   - [ ] Add metrics and statistics commands (INFO)

6. Error Handling
   - [ ] Improve error messages and handling
   - [ ] Add better validation for commands
   - [ ] Implement proper timeout handling

## Contributing

Feel free to submit issues and enhancement requests!

## License

MIT License [here](./LICENSE)
