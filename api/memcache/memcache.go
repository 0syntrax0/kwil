package memcache

import (
	"errors"

	"github.com/google/uuid"
)

var (
	// errors
	ErrUuidExists = errors.New("value already exists")
)

// file info
type File struct {
	Name     string
	Location string
}

// [uuid]FileInfo
type data map[uuid.UUID]File

type Memcache struct {
	Data data
}

// create new instance
func New() *Memcache {
	return &Memcache{
		Data: make(data),
	}
}

// check if a given uuid exists
func (s *Memcache) KeyExists(u uuid.UUID) bool {
	_, found := s.Data[u]
	return found
}
