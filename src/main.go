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

func main() {
	menu.MainMenu()
	pilihanMain := menu.PilihMenu(3)
	adminDefaultData := data.GetAdminData()
	data.AccountDataArr[0] = adminDefaultData
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
					fmt.Println("Akun tak ditemukan!")
				} else {
					switch pilihan {
					case 1:
						mainAdmin()
						pilihan = 4
					case 2:
						// menu pembeli
						mainPembeli()
						pilihan = 4
					case 3:
						// menu penjual
						mainPenjual()
						pilihan = 4
					}
				}
			}
			menu.MainMenu()
			pilihanMain = menu.PilihMenu(3)

		} else {
			var isSuccessAccount, isSuccessStore bool
			isSuccessAccount = false
			isSuccessStore = false
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
					isSuccessAccount = auth.CreateNewAccount(uname, pw, "PEMBELI")
				} else {
					isSuccessAccount = auth.CreateNewAccount(uname, pw, "PENJUAL")
				}
				if isSuccessAccount {
					fmt.Println("Berhasil menambahkan akun, tunggu verifikasi admin")
				} else {
					fmt.Println("Akun gagal ditambahkan")
				}
			}
			for !isSuccessStore && pilihan == 2 {
				storeName := menu.MenuFormRegistrasiToko()
				isSuccessStore = store.CreateNewStore(storeName)
				if isSuccessStore {
					fmt.Println("Berhasil menambahkan toko")
				}
			}
		}
	}
	fmt.Println("Sampai Jumpa Kembali...")
	fmt.Println("=========================================")
	os.Exit(1)
}

func mainPenjual() {
	fmt.Println("Ini menu penjual")
}

func mainPembeli() {
	fmt.Println("Ini menu pembeli")
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
			}
		}
		newData := menu.MenuAdminTindakan(no, accNotVerified, sumAccount)
		admin.UpdateAccountData(newData)
	}
}
