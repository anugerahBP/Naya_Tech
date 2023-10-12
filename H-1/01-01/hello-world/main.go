package main

import "fmt"

func main() {
	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	var firstName = "john"

	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	lastName := "wick"
	lastName = "ethan"
	lastName = "bourne"

	fmt.Printf("halo %s %s!\n", firstName, lastName)
}
