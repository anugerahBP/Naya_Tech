package ruang

import "naya0502/calc"

func Kubus(s int) int {
	return calc.Times(s, calc.Times(s, s))
}
