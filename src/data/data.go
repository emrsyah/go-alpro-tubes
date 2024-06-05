package models

const NMAX int = 100

type AccountData [NMAX]Account
type StoreData [NMAX]Store
type ProductData [NMAX]Product
type TransactionData [NMAX]Transaction
type ProductWithStoreData [NMAX]ProductWithStore

var AccountDataArr AccountData
var StoreDataArr StoreData
var ProductDataArr ProductData
var TransactionDataArr TransactionData

var StoreDataArrLength = 0
var ProductDataArrLength = 0
var TransactionDataArrLength = 0

// Seeding
var seedPenjual = Account{
	Id:         10,
	Username:   "p1",
	Password:   "12345678",
	IsAdmin:    true,
	IsVerified: "accepted",
	Role:       "PENJUAL",
}

var seedPembeli = Account{
	Id:         12,
	Username:   "p2",
	Password:   "12345678",
	IsAdmin:    true,
	IsVerified: "accepted",
	Role:       "PEMBELI",
}

func GetPenjualSeedData() Account {
	return seedPenjual
}

func GetPembeliSeedData() Account {
	return seedPembeli
}

var seedStore = Store{
	Id:     10,
	UserId: 10,
	Name:   "Toko1",
}

func GetStoreSeedData() Store {
	return seedStore
}

var seedBarang1 = Product{
	Id:      10,
	Name:    "Baju",
	Price:   100000,
	Stock:   100,
	StoreId: 10,
}

var seedBarang2 = Product{
	Id:      20,
	Name:    "Celana",
	Price:   200000,
	Stock:   100,
	StoreId: 10,
}

var seedBarang3 = Product{
	Id:      30,
	Name:    "Sepatu",
	Price:   300000,
	Stock:   100,
	StoreId: 10,
}

func GetBarangSeedData() []Product {
	return []Product{seedBarang1, seedBarang2, seedBarang3}
}

var seedTxn1 = Transaction{
	Id:          10,
	AccountId:   10,
	AccountName: "p1",
	StoreId:     10,
	ProductId:   10,
	ProductName: "Baju",
	Quantity:    1,
}

var seedTxn2 = Transaction{
	Id:          20,
	AccountId:   10,
	AccountName: "p1",
	StoreId:     10,
	ProductId:   20,
	ProductName: "Celana",
	Quantity:    6,
}

var seedTxn3 = Transaction{
	Id:          30,
	AccountId:   10,
	AccountName: "p1",
	StoreId:     10,
	ProductId:   30,
	ProductName: "Sepatu",
	Quantity:    5,
}

func GetTxnSeedData() []Transaction {
	return []Transaction{seedTxn1, seedTxn2, seedTxn3}
}

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
