package storage

type LogRecord struct {
	key       string
	value     []byte
	Tombstone bool // true будет означать, что ключ удалён
}

type CommitLog interface {
	Append(record LogRecord) error
	Replay(apply func(LogRecord) error) error
	Truncate() error
	Close() error
}
