package services

import (
	m "github.com/emrsyah/go-alpro-tubes/src/data"
	"github.com/google/uuid"
)

func CheckStoreAvailability(storeName string) bool {
	for i := 0; i <= m.StoreDataArrLength; i++ {
		if m.StoreDataArr[i].Name == storeName {
			return false
		}
	}
	return true
}

func CreateNewStore(name string, uid int64) bool {
	if !CheckStoreAvailability(name) {
		return false
	}
	var store m.Store
	id := uuid.New()
	store.Id = int64(id.ID())
	store.Name = name
	store.UserId = uid
	m.StoreDataArr[m.StoreDataArrLength] = store
	m.StoreDataArrLength += 1
	return true
}

func GetStoreData(ownerId int64) m.Store {
	for i := 0; i < m.StoreDataArrLength; i++ {
		if m.StoreDataArr[i].UserId == (ownerId) {
			return m.StoreDataArr[i]
		}
	}
	return m.Store{}
}

func GetItemDataByStoreId(storeId int64) (m.ProductData, int) {
	var prodArr m.ProductData
	nData := 0
	for i := 0; i < m.ProductDataArrLength; i++ {
		if m.ProductDataArr[i].StoreId == int64(storeId) {
			prodArr[nData] = m.ProductDataArr[i]
			nData++
		}
	}
	return prodArr, nData
}

func GetTxnDataByStoreId(storeId int64) (m.TransactionData, int) {
	var txnArr m.TransactionData
	nData := 0
	for i := 0; i < m.TransactionDataArrLength; i++ {
		if m.TransactionDataArr[i].StoreId == int64(storeId) {
			txnArr[nData] = m.TransactionDataArr[i]
			nData++
		}
	}
	return txnArr, nData
}

func DeleteItemData(itemId int64) {
	isNemu := false
	for i := 0; i < m.ProductDataArrLength; i++ {
		if m.ProductDataArr[i].Id == itemId {
			isNemu = true
			// m.ProductDataArr[i] = m.ProductDataArr[m.ProductDataArrLength-1]
			// m.ProductDataArrLength -= 1
		}
		// Mundurin semua indeks ketika udah nemu yang sama id nya, jadi barang yg dihapus ketimpa sama yang indeks selanjutnya dan seterusnya
		if isNemu {
			m.ProductDataArr[i] = m.ProductDataArr[i+1]
		}
	}
	m.ProductDataArrLength -= 1
}

func EditItemData(itemId int64, newItem m.Product) {
	for i := 0; i < m.ProductDataArrLength; i++ {
		if m.ProductDataArr[i].Id == itemId {
			m.ProductDataArr[i].Name = newItem.Name
			m.ProductDataArr[i].Price = newItem.Price
			m.ProductDataArr[i].Stock = newItem.Stock
			break
		}
	}
}

func CreateItemData(newItem m.Product) {
	var prod m.Product
	id := uuid.New()
	prod.Id = int64(id.ID())
	prod.Name = newItem.Name
	prod.Price = newItem.Price
	prod.Stock = newItem.Stock
	prod.StoreId = newItem.StoreId
	m.ProductDataArr[m.ProductDataArrLength] = prod
	m.ProductDataArrLength += 1
}

func SortAscendingTxnByQuantity(data *m.TransactionData, nData int) {
	for i := 0; i < nData-1; i++ {
		sm := data[i]
		smIdx := i
		for j := i + 1; j < nData; j++ {
			if data[j].Quantity < sm.Quantity {
				sm = data[j]
				smIdx = j
			}
		}
		tmp := data[i]
		data[i] = data[smIdx]
		data[smIdx] = tmp
	}
}

func SortDescendingTxnByQuantity(data *m.TransactionData, nData int) {
	for i := 0; i < nData-1; i++ {
		keyIdx := i + 1
		key := data[keyIdx]
		j := i + 1
		for j > 0 && data[j-1].Quantity < key.Quantity {
			data[j] = data[j-1]
			j--
		}
		data[j] = key
	}
}

func GetStoreById(id int64) m.Store {
	for i := 0; i < m.StoreDataArrLength; i++ {
		if m.StoreDataArr[i].Id == id {
			return m.StoreDataArr[i]
		}
	}
	return m.Store{}
}

func GetAllItemData() (m.ProductWithStoreData, int) {
	var tmp m.ProductWithStoreData
	var i int
	for i = 0; i < m.ProductDataArrLength; i++ {
		store := GetStoreById(m.ProductDataArr[i].StoreId)
		tmp[i] = m.ProductWithStore{
			Product: m.ProductDataArr[i],
			Store:   store,
		}
	}
	return tmp, i
}

func SortAscendingItemByName(data *m.ProductWithStoreData, nData int) {
	for i := 1; i < nData; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].Product.Name > key.Product.Name {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func SortDescendingItemByName(data *m.ProductWithStoreData, nData int) {
	for i := 0; i < nData-1; i++ {
		bg := data[i]
		bgIdx := i
		for j := i + 1; j < nData; j++ {
			if data[j].Product.Name > bg.Product.Name {
				bg = data[j]
				bgIdx = j
			}
		}
		tmp := data[i]
		data[i] = bg
		data[bgIdx] = tmp
	}
}

func CreateTxnData(buyer m.Account, data m.ProductWithStore, qnt int) {
	var txn m.Transaction
	id := uuid.New()
	txn.Id = int64(id.ID())
	txn.ProductId = data.Product.Id
	txn.ProductName = data.Product.Name
	txn.StoreId = data.Store.Id
	txn.Quantity = qnt
	txn.AccountId = buyer.Id
	txn.AccountName = buyer.Username
	m.TransactionDataArr[m.TransactionDataArrLength] = txn
	m.TransactionDataArrLength += 1
}

// urutin data secara ascending
func SortProductData() {
	for i := 1; i < m.ProductDataArrLength; i++ {
		key := m.ProductDataArr[i]
		j := i - 1
		for j >= 0 && m.ProductDataArr[j].Id > key.Id {
			m.ProductDataArr[j+1] = m.ProductDataArr[j]
			j--
		}
		m.ProductDataArr[j+1] = key
	}
}

// ini pake binary search
func UpdateItemStock(itemId int64, qnt int) {
	SortProductData()
	ki := 0
	ka := m.ProductDataArrLength - 1
	mi := (ki + ka) / 2
	for ki <= ka {
		if m.ProductDataArr[mi].Id > itemId {
			ka = mi - 1
		} else if m.ProductDataArr[mi].Id < itemId {
			ki = mi + 1
		} else {
			m.ProductDataArr[mi].Stock -= qnt
			break
		}
		mi = (ki + ka) / 2
	}
}

func CheckoutItem(buyer m.Account, data m.ProductWithStore, qnt int) {
	CreateTxnData(buyer, data, qnt)
	UpdateItemStock(data.Product.Id, qnt)
}
