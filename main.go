package main

import "fmt"
import "calculator/fungsi"

func main() {
	var pilih string
	var angka1, angka2 int

	fmt.Print("masukan angka pertama: ")
	fmt.Scanln(&angka1)
	fmt.Print("masukan angka kedua: ")
	fmt.Scanln(&angka2)

	//pilihan
	fmt.Print("masukan pilihan [*/+-]: ")
	fmt.Scanln(&pilih)

	if pilih == "*" {
		fmt.Println("hasil perkalian:", fungsi.Kali(angka1, angka2))
	} else if pilih == "/" {
		fmt.Println("hasil pembagian:", fungsi.Bagi(angka1, angka2))
	} else if pilih == "+" {
		fmt.Println("hasil penjumlahan:", fungsi.Tambah(angka1, angka2))
	} else if pilih == "-" {
		fmt.Println("hasil pengurangan:", fungsi.Kurang(angka1, angka2))
	} else {
		fmt.Println("error pilih operasi")
	}
}
