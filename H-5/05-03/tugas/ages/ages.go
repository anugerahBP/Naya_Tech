package ages

import "sort"

func TwoOldestAges(ages []int) []int {
	// Urutkan array dalam urutan menurun
	sort.Sort(sort.Reverse(sort.IntSlice(ages)))

	// Kembalikan dua umur tertua
	return []int{ages[1], ages[0]}
}
