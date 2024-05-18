package services

import (
	"fmt"
	// m "github.com/emrsyah/go-alpro-tubes/src/data"
	// "github.com/google/uuid"
)

func PilihMenu(max int) int {
	var pilihan int
	fmt.Print("Masukkan Pilihan: ")
	fmt.Scanln(&pilihan)
	for pilihan < 1 && pilihan > max {
		fmt.Println("Tak ada opsi", pilihan, "masukkan kembali pilihan anda")
		fmt.Print("Masukkan Pilihan: ")
		fmt.Scanln(&pilihan)
	}
	return pilihan
}

func MainMenu() {
	fmt.Println("-------------------------------------")
	fmt.Println("              Main Menu              ")
	fmt.Println("-------------------------------------")
	fmt.Println("1. Login Aplikasi					  ")
	fmt.Println("2. Registrasi Akun		              ")
	fmt.Println("3. Exit Program                      ")
	fmt.Println("-------------------------------------")
}

func MenuLogin() {
	fmt.Println("-------------------------------------")
	fmt.Println("              Login Menu             ")
	fmt.Println("-------------------------------------")
	fmt.Println("1. Login Admin					      ")
	fmt.Println("2. Login Pembeli			      	  ")
	fmt.Println("3. Login Penjual				      ")
	fmt.Println("4. Back to Main Menu                 ")
	fmt.Println("-------------------------------------")
}

func MenuRegistrasi() {
	fmt.Println("-------------------------------------")
	fmt.Println("          Registrasi Menu            ")
	fmt.Println("-------------------------------------")
	fmt.Println("1. Registrasi Pembeli			   	  ")
	fmt.Println("2. Registrasi Penjual			      ")
	fmt.Println("3. Back to Main Menu                 ")
	fmt.Println("-------------------------------------")
}

func MenuFormRegistrasiAkun() (string, string) {
	var uname, pw string
	fmt.Println("-------------------------------------")
	fmt.Println("          Registrasi Akun            ")
	fmt.Println("-------------------------------------")
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&uname)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&pw)
	return uname, pw
}

func MenuFormRegistrasiToko() string {
	var name string
	fmt.Println("-------------------------------------")
	fmt.Println("          Registrasi Toko            ")
	fmt.Println("-------------------------------------")
	fmt.Print("Masukkan Nama Toko: ")
	fmt.Scanln(&name)
	return name
	// store.CreateNewStore(name)
}

func MenuFormLoginAkun() (string, string) {
	var uname, pw string
	fmt.Println("-------------------------------------")
	fmt.Println("             Login Akun              ")
	fmt.Println("-------------------------------------")
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&uname)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&pw)
	return uname, pw
	// isSuccess, role := auth.LoginAccount(uname, pw)
	// fmt.Println(isSuccess, role)
}
