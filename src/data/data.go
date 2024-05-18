package models

const NMAX int = 100

type AccountData [NMAX]Account
type StoreData [NMAX]Store
type ProductData [NMAX]Product
type TransactionData [NMAX]Transaction

var AccountDataArr AccountData

// Define the struct type
type adminData struct {
	uname string
	pw    string
}

// Create a package-level variable
var AdminData = adminData{
	uname: "admin1",
	pw:    "",
}

func GetAdminData() adminData {
	return AdminData
}
