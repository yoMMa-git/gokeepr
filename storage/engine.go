package storage

type StorageEngine interface {
	Put(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
	Scan(from, to string) Iterator
	Close() error
}
