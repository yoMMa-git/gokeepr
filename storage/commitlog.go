package storage

type LogRecord struct {
	key       string
	value     []byte
	Tombstone bool
}

type CommitLog interface {
	Append(record LogRecord) error
	Replay() ([]LogRecord, error) // TODO: Replay(apply func(LogRecord) error) error
	Truncate() error
	Close() error
}
