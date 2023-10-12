package main

import (
	"fmt"
)

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(str string) []Tuple {
	charCount := make(map[rune]int)

	for _, char := range str {
		charCount[char]++
	}

	var result []Tuple
	for char, count := range charCount {
		result = append(result, Tuple{Char: char, Count: count})
	}

	return result
}

func main() {
	str := "Naya  kursus"
	result := OrderedCount(str)

	fmt.Printf("OrderedCount(\"%s\") == []Tuple{", str)
	for i, tuple := range result {
		fmt.Printf("Tuple{'%c', %d}", tuple.Char, tuple.Count)
		if i < len(result)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("}\n")
}
