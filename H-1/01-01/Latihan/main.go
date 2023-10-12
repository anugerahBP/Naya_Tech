package main

import (
	"fmt"
	"strconv"
)

func main() {
	var operator string
	var value1, value2 float64

	// Meminta input operator dari pengguna
	fmt.Print("Masukkan operator aritmatika (+, -, *, /): ")
	fmt.Scanln(&operator)

	// Meminta input nilai variabel pertama dari pengguna
	fmt.Print("Masukkan nilai variabel pertama: ")
	fmt.Scanln(&value1)

	// Meminta input nilai variabel kedua dari pengguna
	fmt.Print("Masukkan nilai variabel kedua: ")
	fmt.Scanln(&value2)

	var result float64

	// Melakukan operasi aritmatika sesuai dengan operator yang dimasukkan
	switch operator {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		result = value1 / value2
	default:
		fmt.Println("Operator yang dimasukkan tidak valid")
		return
	}

	// Menampilkan hasil operasi
	fmt.Println("Hasil:", strconv.FormatFloat(result, 'f', -1, 64))
}
