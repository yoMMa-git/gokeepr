package storage

type Iterator interface {
	Next() bool // наличие следующей записи, если нет - false
	Key() string
	Value() []byte
	Tombstone() bool
}
