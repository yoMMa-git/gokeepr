package storage

type SSTable interface {
	Get(key string) ([]byte, error)
	Scan(from, to string) Iterator
	Size() int64
}

type SSTableWriter interface {
	Add(key string, value []byte, tombstone bool) error
	Finish() (SSTable, error)
}
