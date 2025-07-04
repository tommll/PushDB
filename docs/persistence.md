# Usage

- Client execute commands
```
SET user:1 "Alice"
ZADD leaderboard 100 player1
SET user:2 "Bob"
DEL user:1
```

- WAL logs
```
[CRC32][Size][LSN:1][SET user:1 "Alice"]
[CRC32][Size][LSN:2][ZADD leaderboard 100 player1]
[CRC32][Size][LSN:3][SET user:2 "Bob"]
[CRC32][Size][LSN:4][DEL user:1]
```

- WAL entry format
```
┌─────────────┬─────────────┬─────────────────┐
│ CRC32 (4B)  │ Size (4B)   │ WAL Data (N B)  │
└─────────────┴─────────────┴─────────────────┘
```

**Components**
- CRC32: checksum for data integrity
- Size: length of WAL data
- WAL data: serilized version of the command


### Data flow
- Client sends command
- Log commands to WAL file
- Execute command in memory
- Return response to client

### Data structure
- Write-ahead-log (WAL)

```go
type WAL interface {
    Init() error
    Stop()
    LogCommand(c *Command) error
    ReplayCommand (cb func(*Command) error) error
}

type Element struct {
    Lsn unit64
    ElementType unit8
    Payload []byte
}
```


# Questions

### When reading data, how can we know the data is not corrupted?
- By calculating the checksum of the data and compare with the previously stored checksum

### Why rotate segments?
- To make sure the size of the file is manageable (split data into another file if size in current file reaches maximum segment size)

