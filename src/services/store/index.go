package services

import (
	m "github.com/emrsyah/go-alpro-tubes/src/data"
	"github.com/google/uuid"
)

func CreateNewStore(name string) bool {
	var store m.Store
	id := uuid.New()
	store.Id = int64(id.ID())
	store.Name = name
	return true
}
