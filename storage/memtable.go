package storage

type Memtable interface {
	Get(key string) ([]byte, error)
	Put(key string, value []byte) error
	Delete(key string) error
	Scan(from, to string) Iterator
	Size() int
}
