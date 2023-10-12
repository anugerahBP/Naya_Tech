package main

import "fmt"

func main() {
	fmt.Println("Latihan-1, disini saya memasukkan 1 naga dengan 2 peluru")
	var point = 3

	switch point {
	case 10:
		fmt.Println("Panjang umur")
	case 8, 6, 4, 2:
		fmt.Println("Dapat bertahan hidup")
	default:
		fmt.Println("Pahlawan tersebut mati")
	}
}
