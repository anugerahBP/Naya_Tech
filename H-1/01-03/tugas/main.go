package main

import "fmt"

func summarizeArray(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}

	// Membuat slice baru untuk menyimpan hasil ringkasan
	result := []int{arr[0]}

	for i := 1; i < len(arr); i++ {
		// Mengalikan elemen saat ini dengan elemen sebelumnya
		product := result[len(result)-1] * arr[i]

		// Jika hasil perkalian lebih kecil dari atau sama dengan k, gabungkan elemen
		if product <= k {
			result[len(result)-1] = product
		} else {
			// Jika hasil perkalian lebih besar dari k, tambahkan elemen baru
			result = append(result, arr[i])
		}
	}

	return len(result)
}

func main() {
	// Memasukkan array dan nilai k dari pengguna
	arr := []int{1, 2, 4, 7, 1}
	k := 9

	// Memanggil fungsi summarizeArray untuk meringkas array
	length := summarizeArray(arr, k)

	// Menampilkan hasil
	fmt.Println("Panjang array setelah diringkas:", length)
}
