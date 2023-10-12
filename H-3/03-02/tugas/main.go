package main

import (
	"fmt"
)

var romanToDecimal = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	total := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		value := romanToDecimal[s[i]]

		if value < prevValue {
			total -= value
		} else {
			total += value
		}

		prevValue = value
	}

	return total
}

func main() {
	input1 := "2"
	fmt.Printf("Input: %s\nOutput: %d\n", input1, romanToInt(input1))

	input2 := "II"
	fmt.Printf("Input: %s\nOutput: %d\n", input2, romanToInt(input2))
}
