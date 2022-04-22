package store

import (
	"github.com/PrashantBtkl/distributed-debounce/debouncer/model"
)

type Store interface {
	NewStore(dsn string)
	UpdateBuffer(user int64, bufferDuration uint64) error
	CheckBuffer(user int64, debounceTime int) (*model.DebounceBuffer, error)
}
