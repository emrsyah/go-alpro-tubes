package services

import (
	"fmt"

	d "github.com/emrsyah/go-alpro-tubes/src/data"
	// m "github.com/emrsyah/go-alpro-tubes/src/data"
	// "github.com/google/uuid"
)

// func PilihMenu(max int) int {
// 	var pilihan int
// 	fmt.Println("sebelum", pilihan)
// 	fmt.Print("Masukkan Pilihan: ")
// 	fmt.Scanln(&pilihan)
// 	fmt.Println("sesudah", pilihan)
// 	for pilihan < 1 && pilihan > max {
// 		fmt.Println("Tak ada opsi", pilihan, "masukkan kembali pilihan anda")
// 		fmt.Print("Masukkan Pilihan: ")
// 		fmt.Scanln(&pilihan)
// 	}
// 	return pilihan
// }

func PilihMenu(max int) int {
	var pilihan int
	for {
		fmt.Print("Masukkan Pilihan: ")
		_, err := fmt.Scanf("%d\n", &pilihan)
		// fmt.Print(pilihan)
		// fmt.Print(err)
		if err != nil {
			// if err.Error() == "unexpected newline" {
			// 	continue
			// }
			fmt.Println("Input tidak valid, masukkan angka.")
			continue
		}
		if pilihan >= 1 && pilihan <= max {
			break
		}
		fmt.Println("Tak ada opsi", pilihan, "masukkan kembali pilihan anda")
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

func MenuMainAdminOpen(data d.AccountData, nData int) {
	fmt.Println("-------------------------------------")
	fmt.Println("+            Menu Admin             +")
	fmt.Println("-------------------------------------")
	fmt.Println("+       Data Registrasi Akun        +")
	fmt.Println("-------------------------------------")
	fmt.Println("+ No  +  Nama       +  Role         +")
	fmt.Println("-------------------------------------")

	// Start 1 krn admin idx 0
	for i := 0; i < nData; i++ {
		fmt.Printf("+%4d +  %-10s + %-12s +\n", i+1, data[i].Username, data[i].Role)
		// fmt.Println("-------------------------------------")
	}
	fmt.Println("-------------------------------------")
}

func MenuAdminTindakan(no int, data d.AccountData, nData int) (int64, string) {
	fmt.Println("-------------------------------------")
	fmt.Println("+       Tindakan Pada   Akun        +")
	fmt.Println("-------------------------------------")
	fmt.Printf("+ %-16s + %-14s +\n", data[no].Username, data[no].Role)
	fmt.Println("-------------------------------------")
	fmt.Println("1. Setujui Akun")
	fmt.Println("2. Tolak Akun")
	fmt.Println("3. Kembali")
	pilihan := PilihMenu(3)
	switch pilihan {
	case 1:
		// data[no].IsVerified = "accepted"
		fmt.Println("===================================")
		fmt.Println("Akun berhasil disetujui")
		fmt.Println("===================================")
		return data[no].Id, "accepted"
	case 2:
		data[no].IsVerified = "rejected"
		fmt.Println(data)
		fmt.Println("===================================")
		fmt.Println("Akun berhasil ditolak")
		fmt.Println("===================================")
		return data[no].Id, "rejected"
	case 3:
		return -1, "back"
	default:
		return -1, "back"
	}
}

func MenuMainPenjual(storeName string) int {
	fmt.Println("-------------------------------------")
	fmt.Println("+           Menu Penjual            +")
	fmt.Println("-------------------------------------")
	fmt.Printf("+ %16s + \n", storeName)
	fmt.Println("-------------------------------------")
	fmt.Println("1. Lihat Barang")
	fmt.Println("2. Tambah Barang")
	fmt.Println("3. Lihat Transaksi")
	fmt.Println("4. Kembali")
	pilihan := PilihMenu(4)
	return pilihan
}

func MenuPenjualLihatBarang(data d.ProductData, nData int) int {
	if nData == 0 {
		fmt.Println("==============================")
		fmt.Println("Belum ada barang yang tersedia")
		fmt.Println("==============================")
		return -1
	}
	no := 0

	fmt.Println("---------------------------------------------------")
	fmt.Println("+             Lihat Barang Toko Ini               +")
	fmt.Println("---------------------------------------------------")
	fmt.Println("+ No  +  Nama       +  Harga        +  Stock       +")
	fmt.Println("---------------------------------------------------")
	for i := 0; i < nData; i++ {
		fmt.Printf("+%4d +  %-10s + %12d + %-12d +\n", i+1, data[i].Name, data[i].Price, data[i].Stock)
	}
	fmt.Println("---------------------------------------------------")
	for (no != -1 && no > nData) || (no <= 0 && no != -1) {
		fmt.Print("Masukkan No data yang akan ditindak/ketik -1 untuk kembali: ")
		fmt.Scan(&no)
		fmt.Println()
		if no > nData {
			fmt.Println("Masukkan No data yang benar")
		}
	}
	if no == -1 {
		return -1
	} else {
		return no - 1
	}
}

func MenuPenjualTindakBarang(data d.ProductData, nData int, idxData int) int {
	fmt.Println("-------------------------------------")
	fmt.Println("+       Tindakan Pada Barang        +")
	fmt.Println("-------------------------------------")
	fmt.Println("+ Nama  +  Harga    +  Stok         +")
	fmt.Println("-------------------------------------")
	fmt.Printf("+ %-10s + %-10d + %10d +\n", data[idxData].Name, data[idxData].Price, data[idxData].Stock)
	fmt.Println("-------------------------------------")
	fmt.Println("1. Edit Data Barang")
	fmt.Println("2. Hapus Data Barang")
	fmt.Println("3. Kembali")
	pilihan := PilihMenu(3)
	return pilihan
}

func MenuPenjualTambahDanEditBarang() d.Product {
	var name string
	var price int
	var stock int
	fmt.Print("Masukkan Nama Barang: ")
	fmt.Scanln(&name)
	fmt.Print("Masukkan Harga Barang: ")
	fmt.Scanln(&price)
	fmt.Print("Masukkan Stok Barang: ")
	fmt.Scanln(&stock)
	return d.Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}
}

func MenuPenjualLihatTransaksi(data d.TransactionData, nData int) int {
	if nData == 0 {
		fmt.Println("==============================")
		fmt.Println("Belum ada transaksi yang tersedia")
		fmt.Println("==============================")
		return -1
	}
	fmt.Println("---------------------------------------------------")
	fmt.Println("+             Lihat Transaksi Toko Ini            +")
	fmt.Println("---------------------------------------------------")
	fmt.Println("+ No  +  Nama Barang  +  Pembeli   +  Kuantitas   +")
	fmt.Println("---------------------------------------------------")
	for i := 0; i < nData; i++ {
		fmt.Printf("+%4d +  %-10s + %-12s + %-12d +\n", i+1, data[i].ProductName, data[i].AccountName, data[i].Quantity)
	}
	fmt.Println("---------------------------------------------------")
	fmt.Println("1. Urutkan Ascending Kuantitas")
	fmt.Println("2. Urutkan Descending Kuantitas")
	fmt.Println("3. Kembali")
	pilihan := PilihMenu(3)
	return pilihan
}

func MenuMainPembeli(data d.ProductWithStoreData, nData int) int {
	fmt.Println("-----------------------------------------------------")
	fmt.Println("+                 Menu Pembeli                      +")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("+                 Barang Tersedia                   +")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("+ No  +  Nama Barang  +  Toko   +  Harga   +  Stok  +")
	fmt.Println("-----------------------------------------------------")
	for i := 0; i < nData; i++ {
		fmt.Printf("+%4d +  %-12s + %-7s + %-8d + %-6d +\n", i+1, data[i].Product.Name, data[i].Store.Name, data[i].Price, data[i].Stock)
	}
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1. Beli Barang")
	fmt.Println("2. Urutkan Menaik Berdasar Nama")
	fmt.Println("3. Urutkan Menurun Berdasar Nama")
	fmt.Println("4. Kembali")
	pilihan := PilihMenu(4)
	return pilihan
}

func MenuPembeliPilihBarang(data d.ProductWithStoreData, nData int) int {
	no := 0
	fmt.Println("-----------------------------------------------------")
	fmt.Println("+                 Barang Tersedia                   +")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("+ No  +  Nama Barang  +  Toko   +  Harga   +  Stok  +")
	fmt.Println("-----------------------------------------------------")
	for i := 0; i < nData; i++ {
		fmt.Printf("+%4d +  %-12s + %-7s + %-8d + %-6d +\n", i+1, data[i].Product.Name, data[i].Store.Name, data[i].Price, data[i].Stock)
	}
	fmt.Println("-----------------------------------------------------")
	for (no != -1 && no > nData) || (no <= 0 && no != -1) {
		fmt.Print("Masukkan No data yang akan ditindak/ketik -1 untuk kembali: ")
		fmt.Scan(&no)
		fmt.Println()
		if no > nData {
			fmt.Println("Masukkan No data yang benar")
		}
	}
	if no == -1 {
		return -1
	} else {
		return no - 1
	}
}

func MenuPembeliCheckout(data d.ProductWithStore) int {
	qnt := 0
	fmt.Println("Membeli", data.Product.Name)
	for (qnt <= 0) || (qnt > data.Stock) {
		fmt.Print("Masukkan jumlah barang yang akan dibeli: ")
		fmt.Scan(&qnt)
		if (qnt <= 0) || (qnt > data.Stock) {
			fmt.Println("Jumlah barang yang dibeli harus lebih dari 0 dan tidak lebih dari stock")
		}
	}
	return qnt
}
