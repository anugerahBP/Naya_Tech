package main

import (
	"fmt"
	"strings"
	"unicode"
)

func solve(s string) (string, string) {
	uppercaseCount := 0
	lowercaseCount := 0

	for _, char := range s {
		if unicode.IsUpper(char) {
			uppercaseCount++
		} else {
			lowercaseCount++
		}
	}

	if uppercaseCount > lowercaseCount {
		explanation := fmt.Sprintf("Uppercase characters > lowercase. Change only the \"%c\" to lowercase.", unicode.ToLower(rune(s[lowercaseIndex(s)])))
		return strings.ToUpper(s), explanation
	} else if lowercaseCount > uppercaseCount {
		explanation := fmt.Sprintf("Lowercase characters > uppercase. Change only the \"%c\" to uppercase.", unicode.ToUpper(rune(s[uppercaseIndex(s)])))
		return strings.ToLower(s), explanation
	} else {
		explanation := "Upper == lowercase. Change all to lowercase."
		return strings.ToLower(s), explanation
	}
}

func lowercaseIndex(s string) int {
	for i, char := range s {
		if unicode.IsLower(char) {
			return i
		}
	}
	return -1
}

func uppercaseIndex(s string) int {
	for i, char := range s {
		if unicode.IsUpper(char) {
			return i
		}
	}
	return -1
}

func main() {
	inputs := []string{"Code", "COde", "CoDE"}

	for _, s := range inputs {
		result, explanation := solve(s)
		fmt.Printf("solve(\"%s\") = \"%s\". %s\n", s, result, explanation)
	}
}
