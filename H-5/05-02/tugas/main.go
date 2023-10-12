package main

import "fmt"

func countCuts(width, height int) int {
	if width == 1 && height == 1 {
		return 0
	}
	return (width * height) - 1
}

func main() {
	width1, height1 := 2, 1
	width2, height2 := 3, 1

	cuts1 := countCuts(width1, height1)
	cuts2 := countCuts(width2, height2)

	fmt.Printf("%d x %d -> %d\n", width1, height1, cuts1)
	fmt.Printf("%d x %d -> %d\n", width2, height2, cuts2)
}
