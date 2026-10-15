package storage

type SSTable interface {
	Get(key string) ([]byte, error)
	Scan(from, to string) Iterator
	Size() int64
}
