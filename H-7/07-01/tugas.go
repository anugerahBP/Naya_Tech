package main

import (
	"encoding/json"
	"fmt"
)

// fungsi JSON dari struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{Name: "Demon", Age: 22}
	jsonData, _ := json.Marshal(person) // encode fungsi
	fmt.Println(string(jsonData))

	// input penambahan data
	addPerson("Angel", 28)
}

// fungsi penambahan input data
func addPerson(name string, age int) {
	newPerson := Person{Name: name, Age: age}
	jsonData, _ := json.Marshal(newPerson)
	fmt.Println(string(jsonData))
}
