package services

import (
	m "github.com/emrsyah/go-alpro-tubes/src/data"
	"github.com/google/uuid"
)

func CheckStoreAvailability(storeName string) bool {
	for i := 0; i <= m.StoreDataArrLength; i++ {
		if m.StoreDataArr[i].Name == storeName {
			return false
		}
	}
	return true
}

func CreateNewStore(name string, uid string) bool {
	if !CheckStoreAvailability(name) {
		return false
	}
	var store m.Store
	id := uuid.New()
	store.Id = int64(id.ID())
	store.Name = name
	store.UserId = uid
	m.StoreDataArr[m.StoreDataArrLength] = store
	m.StoreDataArrLength += 1
	return true
}
