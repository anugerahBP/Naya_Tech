package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countNumbers(start, end int) int {
	count := 0

	for i := start; i <= end; i++ {
		if !hasDigit(i, 5) {
			count++
		}
	}

	return count
}

func hasDigit(number, digit int) bool {
	strNumber := strconv.Itoa(number)
	strDigit := strconv.Itoa(digit)
	return strings.Contains(strNumber, strDigit)
}

func main() {
	start1, end1 := 1, 9
	start2, end2 := 4, 17

	result1 := countNumbers(start1, end1)
	result2 := countNumbers(start2, end2)

	fmt.Printf("%d,%d -> %s -> Result %d\n", start1, end1, numbersInRange(start1, end1), result1)
	fmt.Printf("%d,%d -> %s -> Result %d\n", start2, end2, numbersInRange(start2, end2), result2)
}

func numbersInRange(start, end int) string {
	var numbers []string
	for i := start; i <= end; i++ {
		numbers = append(numbers, strconv.Itoa(i))
	}
	return strings.Join(numbers, ",")
}
