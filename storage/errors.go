package storage

import "errors"

var (
	ErrNotFound       = errors.New("key not found")
	ErrDeleted        = errors.New("key deleted") // найден tombstone: искать дальше НЕ нужно
	ErrNotImplemented = errors.New("not implemented")
)
