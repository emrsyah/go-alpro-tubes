package main

import (
	"fmt"
	"os"

	data "github.com/emrsyah/go-alpro-tubes/src/data"
	admin "github.com/emrsyah/go-alpro-tubes/src/services/admin"
	auth "github.com/emrsyah/go-alpro-tubes/src/services/auth"
	menu "github.com/emrsyah/go-alpro-tubes/src/services/menu"
	store "github.com/emrsyah/go-alpro-tubes/src/services/store"
)

func seeding() {
	// seeding data dummy
	adminDefaultData := data.GetAdminData()
	penjualSeedData := data.GetPenjualSeedData()
	pembeliSeedData := data.GetPembeliSeedData()
	data.AccountDataArr[0] = adminDefaultData
	data.AccountDataArr[1] = penjualSeedData
	data.AccountDataArr[2] = pembeliSeedData

	storeSeedData := data.GetStoreSeedData()
	data.StoreDataArr[0] = storeSeedData
	data.StoreDataArrLength = 1

	barangSeedData := data.GetBarangSeedData()
	data.ProductDataArr[0] = barangSeedData[0]
	data.ProductDataArr[1] = barangSeedData[1]
	data.ProductDataArr[2] = barangSeedData[2]
	data.ProductDataArrLength = 3

	txnSeedData := data.GetTxnSeedData()
	data.TransactionDataArr[0] = txnSeedData[0]
	data.TransactionDataArr[1] = txnSeedData[1]
	data.TransactionDataArr[2] = txnSeedData[2]
	data.TransactionDataArrLength = 3
}

func main() {
	seeding()
	menu.MainMenu()
	pilihanMain := menu.PilihMenu(3)
	for pilihanMain != 3 {
		if pilihanMain == 1 {
			// Login Aplikasi
			menu.MenuLogin()
			pilihan := menu.PilihMenu(4)
			for pilihan != 4 {
				// 1 itu admin, 2 pembeli, 3 penjual
				uname, pw := menu.MenuFormLoginAkun()
				isSuccess, _, roleNum := auth.LoginAccount(uname, pw)
				if !isSuccess || roleNum != pilihan {
					if roleNum == -1 || roleNum != pilihan {
						fmt.Println("Akun tak ditemukan!")
					}
				} else {
					accData := auth.GetAccData(uname)
					switch pilihan {
					case 1:
						mainAdmin()
						pilihan = 4
					case 2:
						// menu pembeli
						mainPembeli(accData)
						pilihan = 4
					case 3:
						// menu penjual
						mainPenjual(accData)
						pilihan = 4
					}
				}
			}
			menu.MainMenu()
			pilihanMain = menu.PilihMenu(3)

		} else {
			// Registrasi Akun
			isSuccessAccount := false
			isSuccessStore := false
			uid := int64(-1)
			menu.MenuRegistrasi()
			pilihan := menu.PilihMenu(3)
			if pilihan == 3 {
				menu.MainMenu()
				pilihanMain = menu.PilihMenu(3)
				continue
			}
			// 1 itu regist pembeli, 2 itu regist penjual
			for !isSuccessAccount {
				uname, pw := menu.MenuFormRegistrasiAkun()
				if pilihan == 1 {
					isSuccessAccount, uid = auth.CreateNewAccount(uname, pw, "PEMBELI")
				} else {
					isSuccessAccount, uid = auth.CreateNewAccount(uname, pw, "PENJUAL")
				}
				if isSuccessAccount {
					fmt.Println("Berhasil menambahkan akun, tunggu verifikasi admin")
				} else {
					fmt.Println("Akun gagal ditambahkan")
				}
			}
			for !isSuccessStore && pilihan == 2 {
				storeName := menu.MenuFormRegistrasiToko()
				isSuccessStore = store.CreateNewStore(storeName, uid)
				if isSuccessStore {
					fmt.Println("Berhasil menambahkan toko")
				} else {
					fmt.Println("Gagal menambahkan toko, Coba Nama Toko Lain")
				}
			}
		}
	}
	fmt.Println("Sampai Jumpa Kembali...")
	fmt.Println("=========================================")
	os.Exit(1)
}

func mainPenjual(accData data.Account) {
	storeData := store.GetStoreData(accData.Id)
	itemData, nItem := store.GetItemDataByStoreId(storeData.Id)
	txnData, nTxn := store.GetTxnDataByStoreId(storeData.Id)
	isOnLoop := true
	for isOnLoop {
		pilihan := menu.MenuMainPenjual(storeData.Name)
		switch pilihan {
		case 1:
			// fmt.Println("Lihat Barang")
			// Get Data Barang Toko Ini
			//  Kirim ke menu buat di loop
			// di dalem menu bisa milih dan lakuin tindakan
			indeksBarangDitindak := menu.MenuPenjualLihatBarang(itemData, nItem)
			if indeksBarangDitindak == -1 {
				isOnLoop = false
				continue
			}
			nomorTindakan := menu.MenuPenjualTindakBarang(itemData, nItem, indeksBarangDitindak)
			if nomorTindakan == 3 {
				continue
			} else if nomorTindakan == 1 {
				// Edit Barang
				newItem := menu.MenuPenjualTambahDanEditBarang()
				store.EditItemData(itemData[indeksBarangDitindak].Id, newItem)
				fmt.Println("===========================")
				fmt.Println("  Barang berhasil diedit  ")
				fmt.Println("===========================")
				itemData, nItem = store.GetItemDataByStoreId(storeData.Id)
			} else if nomorTindakan == 2 {
				// Hapus Barang
				store.DeleteItemData(itemData[indeksBarangDitindak].Id)
				fmt.Println("===========================")
				fmt.Println("  Barang berhasil dihapus  ")
				fmt.Println("===========================")
				itemData, nItem = store.GetItemDataByStoreId(storeData.Id)
			}
		case 2:
			// tambah barang ke toko ini
			newItem := menu.MenuPenjualTambahDanEditBarang()
			newItem.StoreId = storeData.Id
			store.CreateItemData(newItem)
			fmt.Println("=============================")
			fmt.Println(" Barang berhasil ditambahkan ")
			fmt.Println("=============================")
			itemData, nItem = store.GetItemDataByStoreId(storeData.Id)
		case 3:
			// lihat semua transaksi toko ini
			var tindakan int
			for tindakan != 3 && tindakan != -1 {
				tindakan = menu.MenuPenjualLihatTransaksi(txnData, nTxn)
				if tindakan == 3 || tindakan == -1 {
					continue
				} else if tindakan == 1 {
					// Urutin Ascending Kuantitas with selection
					fmt.Println("=============================")
					fmt.Println(" Sorting Product Ascending   ")
					fmt.Println("=============================")
					store.SortAscendingTxnByQuantity(&txnData, nTxn)
				} else if tindakan == 2 {
					// Urutin Descending Kuantitas with insertion
					fmt.Println("=============================")
					fmt.Println(" Sorting Product Descending  ")
					fmt.Println("=============================")
					store.SortDescendingTxnByQuantity(&txnData, nTxn)
				}
			}
		case 4:
			fmt.Println("Back to Main Menu")
			isOnLoop = false
		}
	}
}

func mainPembeli(accData data.Account) {
	allItemData, nItemData := store.GetAllItemData()
	isOnLoop := true
	for isOnLoop {
		pilihan := menu.MenuMainPembeli(allItemData, nItemData)
		switch pilihan {
		case 1:
			// Beli Barang
			idxPilihanBarang := menu.MenuPembeliPilihBarang(allItemData, nItemData)
			if idxPilihanBarang == -1 {
				continue
			} else {
				quantityPembelian := menu.MenuPembeliCheckout(allItemData[idxPilihanBarang])
				store.CheckoutItem(accData, allItemData[idxPilihanBarang], quantityPembelian)
				fmt.Println("========================")
				fmt.Println("Barang berhasil dibeli")
				fmt.Println("========================")
				allItemData, nItemData = store.GetAllItemData()
			}
		case 2:
			// Urutkan Ascending Nama
			store.SortAscendingItemByName(&allItemData, nItemData)
		case 3:
			// Urutkan Descending Nama
			store.SortDescendingItemByName(&allItemData, nItemData)
		case 4:
			isOnLoop = false
		}
	}
}

func mainAdmin() {
	isOnLoop := true
	for isOnLoop {
		var no int
		fmt.Println()
		accNotVerified, sumAccount := admin.GetAccountNotVerified()
		if sumAccount == 0 {
			fmt.Println("Belum ada akun yang mendaftar")
			fmt.Println("1. Back to Main Menu 		  ")
			menu.PilihMenu(1)
			isOnLoop = false
			continue
		}
		// fmt.Println(accNotVerified)
		menu.MenuMainAdminOpen(accNotVerified, sumAccount)
		for no == 0 || no > sumAccount {
			fmt.Print("Masukkan No data yang akan ditindak/ketik -1 untuk kembali: ")
			fmt.Scan(&no)
			if no == 0 || no > sumAccount {
				fmt.Println("Masukkan No data yang benar")
			}
			if no == -1 {
				isOnLoop = false
				continue
			}
		}
		if no != -1 {
			uid, status := menu.MenuAdminTindakan(no-1, accNotVerified, sumAccount)
			if status != "back" {
				admin.UpdateAccountData(uid, status)
			}
		}
	}
}
