package models

type Account struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type Store struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	Id      int64   `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	StoreId int64   `json:"store_id"`
}

type Transaction struct {
	Id        int64 `json:"id"`
	AccountId int64 `json:"account_id"`
	StoreId   int64 `json:"store_id"`
	ProductId int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

const NMAX int = 100

// DATA

type AccountData [NMAX]Account
type StoreData [NMAX]Store
type ProductData [NMAX]Product
type TransactionData [NMAX]Transaction
