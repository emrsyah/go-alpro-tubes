package services

import (
	"fmt"

	data "github.com/emrsyah/go-alpro-tubes/src/data"
)

func GetAccountNotVerified() (data.AccountData, int) {
	nData := 0
	var tmpArr data.AccountData
	// fmt.Println(data.AccountDataArr)
	for i := 1; i < len(data.AccountDataArr)-1; i++ {
		if data.AccountDataArr[i].Username != "" && data.AccountDataArr[i].Password != "" && data.AccountDataArr[i].IsVerified == "pending" {
			// fmt.Println("Ada data: ", data.AccountDataArr[i], "idx: ", i)
			tmpArr[nData] = data.AccountDataArr[i]
			nData++
		} else if data.AccountDataArr[i].Username == "" && data.AccountDataArr[i].Password == "" {
			return tmpArr, nData
		}
	}
	// fmt.Println(tmpArr)
	return tmpArr, nData
}

func UpdateAccountData(id int64, status string) {
	for i := 1; len(data.AccountDataArr)-1 > i; i++ {
		if data.AccountDataArr[i].Id == int64(id) {
			data.AccountDataArr[i].IsVerified = status
			break
		}
	}
	fmt.Println("Data Terbaru:", data.AccountDataArr)
}
