package multiples

func ClosestMultipleOf5(num int) int {
	remainder := num % 5
	if remainder >= 3 {
		return num + (5 - remainder)
	}
	return num - remainder
}
