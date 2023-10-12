package main

import "fmt"

func addDigits(num int) int {
	explanation := fmt.Sprintf("Explanation: The process is\n%d", num)
	for num >= 10 {
		sum := 0
		digits := ""
		for num > 0 {
			digit := num % 10
			sum += digit
			digits = fmt.Sprintf("%d + %s", digit, digits)
			num /= 10
		}
		num = sum
		explanation += fmt.Sprintf(" --> %s --> %d", digits[:len(digits)-3], sum)
	}
	explanation += fmt.Sprintf("\nSince %d has only one digit, return it.", num)
	fmt.Println(explanation)
	return num
}

func main() {
	input := 38
	fmt.Printf("Input: %d\nOutput: %d\n", input, addDigits(input))
}
