package main

import "fmt"

func solution(a, b string) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}

	lastCharA := a[len(a)-1]
	lastCharB := b[len(b)-1]

	return lastCharA == lastCharB
}

func main() {
	fmt.Println(solution("abc", "bc")) // Output: true
	fmt.Println(solution("abc", "d"))  // Output: false
}
