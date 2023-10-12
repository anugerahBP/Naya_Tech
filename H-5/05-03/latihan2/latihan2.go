package main

import (
	"fmt"
	"latihan2/multiples"
)

func main() {
	inputs := []int{0, 2, 3, 12, 21, 30, -2, -5}

	fmt.Println("input      output")
	for _, num := range inputs {
		result := multiples.ClosestMultipleOf5(num)
		fmt.Printf("%d     ->     %d\n", num, result)
	}
}
