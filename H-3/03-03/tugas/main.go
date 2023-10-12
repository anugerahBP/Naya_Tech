package main

import (
	"fmt"
	"strings"
)

func plusOne(digits []int) []int {
	n := len(digits)

	// Iterasi dari kanan ke kiri pada array digits
	for i := n - 1; i >= 0; i-- {
		// Jika digit bukan 9, tambahkan 1 dan kembalikan hasilnya
		if digits[i] != 9 {
			digits[i]++
			return digits
		}

		// Jika digit adalah 9, ubah menjadi 0 dan lanjutkan ke digit sebelumnya
		digits[i] = 0
	}

	// Jika masih ada carry yang tersisa (semua digit awal adalah 9), tambahkan angka 1 di depan array
	return append([]int{1}, digits...)
}

func formatOutput(digits []int) string {
	strDigits := make([]string, len(digits))
	for i, digit := range digits {
		strDigits[i] = fmt.Sprintf("%d", digit)
	}
	return fmt.Sprintf("[%s]", strings.Join(strDigits, ","))
}

func main() {
	input1 := []int{1, 2, 3}
	fmt.Println("Input: digits =", formatOutput(input1))
	fmt.Println("Output:", formatOutput(plusOne(input1)))

	input2 := []int{9}
	fmt.Println("Input: digits =", formatOutput(input2))
	fmt.Println("Output:", formatOutput(plusOne(input2)))

	input3 := []int{4, 12}
	fmt.Println("Input: digits =", formatOutput(input3))
	fmt.Println("Output:", formatOutput(plusOne(input3)))
}
