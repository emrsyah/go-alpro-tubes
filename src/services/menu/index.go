package services

import (
	"fmt"

	// m "github.com/emrsyah/go-alpro-tubes/src/data"
	auth "github.com/emrsyah/go-alpro-tubes/src/services/auth"
	store "github.com/emrsyah/go-alpro-tubes/src/services/store"
	// "github.com/google/uuid"
)

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

func MenuFormRegistrasiAkun(role string) {
	var uname, pw string
	fmt.Println("-------------------------------------")
	fmt.Println("          Registrasi Akun            ")
	fmt.Println("-------------------------------------")
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&uname)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&pw)
	auth.CreateNewAccount(uname, pw, role)
}

func MenuFormRegistrasiToko(usename string) {
	var name string
	fmt.Println("-------------------------------------")
	fmt.Println("          Registrasi Toko            ")
	fmt.Println("-------------------------------------")
	fmt.Print("Masukkan Nama Toko: ")
	fmt.Scanln(&name)
	store.CreateNewStore(name)
}

func MenuFormLoginAkun() {
	var uname, pw string
	fmt.Println("-------------------------------------")
	fmt.Println("             Login Akun              ")
	fmt.Println("-------------------------------------")
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&uname)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&pw)
	isSuccess, role := auth.LoginAccount(uname, pw)
	fmt.Println(isSuccess, role)
}
