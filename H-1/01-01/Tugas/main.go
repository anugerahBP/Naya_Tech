package main

import (
	"fmt"
	"math"
)

func main() {
	var pilihan int

	// Menampilkan menu pilihan kepada pengguna
	fmt.Println("Menu Pilihan:")
	fmt.Println("1. Persegi")
	fmt.Println("2. Lingkaran")
	fmt.Println("3. Kubus")
	fmt.Println("4. Tabung")

	fmt.Print("Pilih bangun datar/ruang: ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		// Menghitung keliling dan luas persegi
		var sisi float64

		fmt.Print("Masukkan panjang sisi persegi: ")
		fmt.Scanln(&sisi)

		keliling := 4 * sisi
		luas := sisi * sisi

		fmt.Println("Keliling persegi:", keliling)
		fmt.Println("Luas persegi:", luas)

	case 2:
		// Menghitung keliling dan luas lingkaran
		var jariJari float64

		fmt.Print("Masukkan panjang jari-jari lingkaran: ")
		fmt.Scanln(&jariJari)

		keliling := 2 * math.Pi * jariJari
		luas := math.Pi * jariJari * jariJari

		fmt.Println("Keliling lingkaran:", keliling)
		fmt.Println("Luas lingkaran:", luas)

	case 3:
		// Menghitung keliling dan volume kubus
		var sisi float64

		fmt.Print("Masukkan panjang sisi kubus: ")
		fmt.Scanln(&sisi)

		keliling := 12 * sisi
		volume := math.Pow(sisi, 3)

		fmt.Println("Keliling kubus:", keliling)
		fmt.Println("Volume kubus:", volume)

	case 4:
		// Menghitung keliling dan volume tabung
		var jariJari, tinggi float64

		fmt.Print("Masukkan panjang jari-jari tabung: ")
		fmt.Scanln(&jariJari)

		fmt.Print("Masukkan tinggi tabung: ")
		fmt.Scanln(&tinggi)

		keliling := 2 * math.Pi * jariJari
		volume := math.Pi * jariJari * jariJari * tinggi

		fmt.Println("Keliling tabung:", keliling)
		fmt.Println("Volume tabung:", volume)

	default:
		fmt.Println("Pilihan tidak valid")
	}
}
