package main

import "fmt"

func main() {
	var pilih int
	for {
		menu()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			hitungJumlah()
		case 2:
			hitungKali()
		case 3:
			hitungBagi()
		case 4:
			fmt.Println("Terima kasih telah menggunakan kalkulator.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		}
	}
}

func menu() {
	fmt.Println("-----------------------")
	fmt.Println("        M E N U")
	fmt.Println("-----------------------")
	fmt.Println("1. Hitung Penjumlahan")
	fmt.Println("2. Hitung Perkalian")
	fmt.Println("3. Hitung Pembagian")
	fmt.Println("4. Exit")
	fmt.Println("-----------------------")
}

func hitungJumlah() {
	var a, b int
	fmt.Print("Masukkan dua bilangan yang akan dijumlahkan: ")
	fmt.Scanln(&a, &b)
	fmt.Printf("Hasil penjumlahan: %d\n", a+b)
}

func hitungKali() {
	var a, b int
	fmt.Print("Masukkan dua bilangan yang akan dikalikan: ")
	fmt.Scanln(&a, &b)
	fmt.Printf("Hasil perkalian: %d\n", a*b)
}

func hitungBagi() {
	var a, b float32
	fmt.Print("Masukkan dua bilangan yang akan dibagikan: ")
	fmt.Scanln(&a, &b)
	if b != 0 {
		fmt.Printf("Hasil pembagian: %1.1f\n", a/b)
	} else {
		fmt.Println("Tidak dapat melakukan pembagian oleh nol.")
	}
}
