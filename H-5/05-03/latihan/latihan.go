package main

import (
	"fmt"
	"latihan/middleindex"
)

func main() {
	arr1 := []int{2, 3, 1}
	arr2 := []int{5, 10, 14}

	index1 := middleindex.Gimme(arr1)
	index2 := middleindex.Gimme(arr2)

	fmt.Printf("gimme(%v) => %d\n", arr1, index1)
	fmt.Printf("gimme(%v) => %d\n", arr2, index2)
}
