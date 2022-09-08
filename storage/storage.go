package storage

// Storage interface
type Storage interface {
	Init() error
	Create(key []byte, val []byte)
	Read(key []byte) []byte
	Close() error
}
