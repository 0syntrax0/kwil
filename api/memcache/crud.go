package memcache

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

// mimic SQL's insert
func (s *Memcache) Insert(u uuid.UUID, f File) error {
	if s.KeyExists(u) {
		slog.Error("unable to create NEW record", "uuid", u.String(), "error", ErrUuidExists.Error())
		return ErrUuidExists
	}

	s.Data[u] = f
	return nil
}

// mimic SQL's select
func (s *Memcache) Select(u uuid.UUID) (File, error) {
	if !s.KeyExists(u) {
		slog.Warn("key does not exist", "uuid", u.String())
		return File{}, fmt.Errorf("%v not found", u)
	}

	return s.Data[u], nil
}

func (s *Memcache) SelectAll() data {
	return s.Data
}

// mimic SQL's delte
func (s *Memcache) Delete(u uuid.UUID) error {
	if !s.KeyExists(u) {
		slog.Error("unable to delete", "uuid", u.String())
		return fmt.Errorf("%v not found", u)
	}

	delete(s.Data, u)
	return nil
}

// mimic SQL's update
func (s *Memcache) Update(u uuid.UUID, f File) error {
	if !s.KeyExists(u) {
		slog.Error("unable to update", "uuid", u.String(), "file-name", f.Name, "file-location", f.Location)
		return fmt.Errorf("%v not found", u)
	}

	s.Data[u] = f
	return nil
}
