package main

import "fmt"

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
		fmt.Println("hasil perkalian:", kali(angka1, angka2))
	} else if pilih == "/" {
		fmt.Println("hasil pembagian:", bagi(angka1, angka2))
	} else if pilih == "+" {
		fmt.Println("hasil penjumlahan:", tambah(angka1, angka2))
	} else if pilih == "-" {
		fmt.Println("hasil pengurangan:", kurang(angka1, angka2))
	} else {
		fmt.Println("error pilih operasi")
	}
}

func kali(a, b int) int {
	return a * b
}

func bagi(a, b int) int {
	return a / b

}

func tambah(a, b int) int {
	return a + b
}

func kurang(a, b int) int {
	return a - b
}
