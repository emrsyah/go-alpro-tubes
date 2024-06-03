package service

import (
	"fmt"

	m "github.com/emrsyah/go-alpro-tubes/src/data"
	"github.com/google/uuid"
)

func CreateNewAccount(uname string, pw string, role string) (bool, int64) {
	var acc m.Account
	id := uuid.New()
	acc.Id = int64(id.ID())
	acc.Username = uname
	acc.Password = pw
	acc.IsAdmin = false
	acc.IsVerified = "pending"
	acc.Role = role
	nData := 0
	for i := 0; i < len(m.AccountDataArr)-1; i++ {
		if m.AccountDataArr[i].Username == uname {
			return false, -1
		}
		if m.AccountDataArr[i].Username != "" && m.AccountDataArr[i].Password != "" {
			nData++
		}
	}
	m.AccountDataArr[nData] = acc
	// fmt.Println(m.AccountDataArr)
	return true, acc.Id
}

func LoginAccount(uname string, pw string) (bool, string, int) {
	// car dan lihat role
	// fmt.Println(m.AccountDataArr)
	for i := 0; i < len(m.AccountDataArr); i++ {
		// fmt.Println(i)
		// fmt.Println(m.AccountDataArr[i].Username)
		if m.AccountDataArr[i].Username == uname && m.AccountDataArr[i].Password == pw {
			var role int
			if m.AccountDataArr[i].Role == "ADMIN" {
				role = 1
			} else if m.AccountDataArr[i].Role == "PEMBELI" {
				role = 2
			} else {
				role = 3
			}
			if m.AccountDataArr[i].IsVerified == "pending" {
				fmt.Println("Akun belum diverifikasi Admin")
				return false, m.AccountDataArr[i].Role, role
			}
			if m.AccountDataArr[i].IsVerified == "rejected" {
				fmt.Println("Akun ditolak Admin")
				return false, m.AccountDataArr[i].Role, role
			}
			return true, m.AccountDataArr[i].Role, role
		}
		if m.AccountDataArr[i].Username == "" && m.AccountDataArr[i].Password == "" {
			return false, "-", -1
		}
	}
	return false, "-", -1
}

func GetAccData(uname string) m.Account {
	for i := 0; i < len(m.AccountDataArr); i++ {
		if m.AccountDataArr[i].Username == uname {
			return m.AccountDataArr[i]
		}
	}
	return m.Account{}
}
