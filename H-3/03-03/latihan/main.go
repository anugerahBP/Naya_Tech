package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	carry := 0
	result := ""

	i := len(a) - 1
	j := len(b) - 1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		result = fmt.Sprintf("%d%s", sum%2, result)
		carry = sum / 2
	}

	return result
}

func main() {
	input1_a := "11"
	input1_b := "1"
	fmt.Printf("Input: a = %s, b = %s\nOutput: %s\n", input1_a, input1_b, addBinary(input1_a, input1_b))

	input2_a := "1010"
	input2_b := "1011"
	fmt.Printf("Input: a = %s, b = %s\nOutput: %s\n", input2_a, input2_b, addBinary(input2_a, input2_b))
}
