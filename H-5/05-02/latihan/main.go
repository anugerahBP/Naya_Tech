package main

import (
	"fmt"
	"strconv"
	"strings"
)

func minMax(arr []int) (int, int) {
	min := arr[0]
	max := arr[0]

	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{2334454, 5}
	arr3 := []int{1}

	min1, max1 := minMax(arr1)
	min2, max2 := minMax(arr2)
	min3, max3 := minMax(arr3)

	output1 := fmt.Sprintf("[%d,%d]", min1, max1)
	output2 := fmt.Sprintf("[%d,%d]", min2, max2)
	output3 := fmt.Sprintf("[%d,%d]", min3, max3)

	fmt.Printf("%s --> %s\n", arrToString(arr1), output1)
	fmt.Printf("%s --> %s\n", arrToString(arr2), output2)
	fmt.Printf("%s --> %s\n", arrToString(arr3), output3)
}

func arrToString(arr []int) string {
	strArr := make([]string, len(arr))
	for i, num := range arr {
		strArr[i] = strconv.Itoa(num)
	}
	return "[" + strings.Join(strArr, ",") + "]"
}
