package models

type Account struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
	// IsVerified bisa pending, success, rejected
	IsVerified string `json:"is_verified"`
	Role       string `json:"role"`
}

type Store struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"user_id"`
	Name   string `json:"name"`
}

type Product struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Stock   int    `json:"stock"`
	StoreId int64  `json:"store_id"`
}

type Transaction struct {
	Id          int64  `json:"id"`
	AccountId   int64  `json:"account_id"`
	AccountName string `json:"account_name"`
	StoreId     int64  `json:"store_id"`
	ProductId   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}
