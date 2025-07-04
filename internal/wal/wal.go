package wal

type WAL interface {
	Init() error
	Stop() error
	LogCommand(command string) error
}

type Element struct {
	Lsn         uint64
	ElementType uint8
	Payload     []byte
}
