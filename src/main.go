package main

import (
	"fmt"
	"os"

	auth "github.com/emrsyah/go-alpro-tubes/src/services/auth"
	menu "github.com/emrsyah/go-alpro-tubes/src/services/menu"
	store "github.com/emrsyah/go-alpro-tubes/src/services/store"
)

func main() {
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
					fmt.Println("Akun tak ditemukan!")
				}
				switch pilihan {
				case 1:
					mainAdmin()
				case 2:
					// menu pembeli
					mainPembeli()
				case 3:
					// menu penjual
					mainPenjual()
				}
			}
			continue
		} else {
			// Registrasi Akun
			var isSuccessAccount, isSuccessStore bool
			isSuccessAccount = false
			isSuccessStore = false
			menu.MenuRegistrasi()
			pilihan := menu.PilihMenu(3)
			if pilihan == 3 {
				continue
			}
			// 1 itu regist pembeli, 2 itu regist penjual
			for !isSuccessAccount {
				uname, pw := menu.MenuFormRegistrasiAkun()
				if pilihan == 1 {
					isSuccessAccount = auth.CreateNewAccount(uname, pw, "Pembeli")
				} else {
					isSuccessAccount = auth.CreateNewAccount(uname, pw, "Penjual")
				}
				if isSuccessAccount {
					fmt.Println("Berhasil menambahkan akun")
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

}

func mainPembeli() {

}

func mainAdmin() {

}
