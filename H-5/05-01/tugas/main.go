package main

import (
	"fmt"
)

// Struct Bus untuk menyimpan data jumlah penumpang dan total penumpang
type Bus struct {
	totalPenumpang int
}

// Method untuk menambahkan jumlah penumpang di pemberhentian
func (b *Bus) TambahPenumpang(naik, turun int) {
	b.totalPenumpang += naik - turun
}

// Factory function untuk membuat objek Bus baru
func NewBus() *Bus {
	return &Bus{}
}

func main() {
	// Membuat objek Bus baru
	bus := NewBus()

	// Array yang berisi array penumpang naik dan turun di setiap pemberhentian
	pemberhentian := [][]int{
		{5, 0},
		{3, 2},
		{2, 1},
		{0, 3},
	}

	// Menghitung jumlah penumpang di setiap pemberhentian
	for _, p := range pemberhentian {
		bus.TambahPenumpang(p[0], p[1])
	}

	// Menampilkan jumlah penumpang terakhir di pemberhentian terakhir
	fmt.Printf("Jumlah penumpang terakhir di pemberhentian terakhir: %d\n", bus.totalPenumpang)
}
