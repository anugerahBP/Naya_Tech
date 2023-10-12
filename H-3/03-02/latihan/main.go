package main

import (
	"fmt"
)

func mxdiflg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	minA1, maxA1 := len(a1[0]), len(a1[0])
	minA2, maxA2 := len(a2[0]), len(a2[0])

	// Cari panjang terpendek dan terpanjang dari masing-masing array
	for _, str := range a1 {
		length := len(str)
		if length < minA1 {
			minA1 = length
		}
		if length > maxA1 {
			maxA1 = length
		}
	}

	for _, str := range a2 {
		length := len(str)
		if length < minA2 {
			minA2 = length
		}
		if length > maxA2 {
			maxA2 = length
		}
	}

	// Hitung jarak terjauh antara perbandingan elemen terpendek dan terpanjang
	return max(maxA1-minA2, maxA2-minA1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	a2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}

	result := mxdiflg(a1, a2)
	fmt.Println("Jarak terjauh dari perbandingan elemen terpendek dan terpanjang:", result)
}
