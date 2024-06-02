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
	UserId string `json:"user_id"`
	Name   string `json:"name"`
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
