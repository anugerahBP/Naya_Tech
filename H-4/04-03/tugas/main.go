package main

import "fmt"

func isTriangle(a, b, c int) bool {
	// Pemeriksaan ketentuan segitiga
	if a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}

/*
Buatlah sebuah program yang menentukan apakah kita dapat membuat sebuah segitia
dengan parameter yang di dapat (a, b, c).
Jika iya kembalikan nilai true dan false jika tidak.
*/
func main() {
	a, b, c := 3, 4, 5

	if isTriangle(a, b, c) {
		/* Hasil True */
		fmt.Println("Dapat membuat segitiga dengan panjang sisi:", a, b, c)
	} else {
		/* Hasil False */
		fmt.Println("Tidak dapat membuat segitiga dengan panjang sisi:", a, b, c)
	}
}
