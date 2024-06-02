package models

const NMAX int = 100

type AccountData [NMAX]Account
type StoreData [NMAX]Store
type ProductData [NMAX]Product
type TransactionData [NMAX]Transaction

var AccountDataArr AccountData
var StoreDataArr StoreData
var StoreDataArrLength = 0
var ProductDataArr ProductData
var TransactionDataArr TransactionData

// Define the struct type

// Create a package-level variable
var AdminData = Account{
	Id:         1,
	Username:   "ADMIN1",
	Password:   "12345678",
	IsAdmin:    true,
	IsVerified: "accepted",
	Role:       "ADMIN",
}

func GetAdminData() Account {
	return AdminData
}
