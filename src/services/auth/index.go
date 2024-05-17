package service

import (
	m "github.com/emrsyah/go-alpro-tubes/src/data"
	"github.com/google/uuid"
)

func CreateNewAccount(uname string, pw string, role string) {
	var acc m.Account
	id := uuid.New()
	id = uuid.New()
	acc.Id = int64(id.ID())
	acc.Username = uname
	acc.Password = pw
	acc.IsAdmin = false
	acc.IsVerified = false
	acc.Role = role
}

func LoginAccount(uname string, pw string) (bool, string) {
	// car dan lihat role

	return true, "galse"
}
