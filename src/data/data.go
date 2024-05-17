package models

const NMAX int = 100

type AccountData [NMAX]Account
type StoreData [NMAX]Store
type ProductData [NMAX]Product
type TransactionData [NMAX]Transaction

var accountData AccountData
