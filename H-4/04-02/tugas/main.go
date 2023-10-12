package main

import "fmt"

func isUgly(n int) (bool, string) {
	if n <= 0 {
		return false, ""
	}

	primeFactors := []int{2, 3, 5}
	explanation := ""

	for _, factor := range primeFactors {
		for n%factor == 0 {
			n /= factor
			if explanation != "" {
				explanation += " × "
			}
			explanation += fmt.Sprintf("%d", factor)
		}
	}

	if n == 1 {
		return true, explanation
	}

	return false, ""
}

func main() {
	input1 := 6
	isUgly1, explanation1 := isUgly(input1)
	fmt.Printf("Input: n = %d\nOutput: %t\n", input1, isUgly1)
	if explanation1 != "" {
		fmt.Printf("Explanation: %d = %s\n", input1, explanation1)
	}

	input2 := 56
	isUgly2, explanation2 := isUgly(input2)
	fmt.Printf("Input: n = %d\nOutput: %t\n", input2, isUgly2)
	if explanation2 != "" {
		fmt.Printf("Explanation: %d = %s\n", input2, explanation2)
	}
}
