package services

import (
	data "github.com/emrsyah/go-alpro-tubes/src/data"
)

func GetAccountNotVerified() (data.AccountData, int) {
	nData := 0
	var tmpArr data.AccountData
	for i := 1; i < len(data.AccountDataArr)-1; i++ {
		if data.AccountDataArr[i].Username != "" && data.AccountDataArr[i].Password != "" && data.AccountDataArr[i].IsVerified == "pending" {
			tmpArr[i] = data.AccountDataArr[i]
			nData++
		} else if data.AccountDataArr[i].Username == "" && data.AccountDataArr[i].Password == "" {
			return tmpArr, nData
		}
	}
	return tmpArr, nData
}

func UpdateAccountData(newData data.AccountData) {
	data.AccountDataArr = newData
}
