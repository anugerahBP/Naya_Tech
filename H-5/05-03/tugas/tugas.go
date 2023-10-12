package main

import (
	"fmt"
	"tugas/ages"
)

func main() {
	ages1 := []int{1, 2, 10, 8}
	ages2 := []int{1, 5, 87, 45, 8, 8}
	ages3 := []int{1, 3, 10, 0}

	result1 := ages.TwoOldestAges(ages1)
	result2 := ages.TwoOldestAges(ages2)
	result3 := ages.TwoOldestAges(ages3)

	fmt.Printf("%v --> %v\n", ages1, result1)
	fmt.Printf("%v --> %v\n", ages2, result2)
	fmt.Printf("%v --> %v\n", ages3, result3)
}
