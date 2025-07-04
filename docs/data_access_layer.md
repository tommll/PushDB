# 1. Data types
- String

### Data structure
```go
// Core objects: store the actual data values

type Obj struct {
    Type uint8
    Value interface{}
}

// Storage: handle persistence and data retrieval
type Store struct {
    store map[string, *Obj]
    expires map[*Obj, int64]
    evictionStrategy EvictionStrategy
}

// Serialization for sending data over the network
func serialize(obj *Obj) ([]byte, error)
func deSerialize(data []byte) (obj *Obj, error)

```


