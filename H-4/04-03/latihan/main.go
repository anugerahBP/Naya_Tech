package main

import "fmt"

func nbYear(p0, percent, aug, p int) int {
	population := p0
	years := 0

	for population < p {
		increase := int(float64(population) * float64(percent) / 100)
		population += increase + aug
		years++

		fmt.Printf("At the end of the %d year there will be:\n%d + %d * %.2f + %d => %d inhabitants\n",
			years, p0, p0, float64(percent)/100, aug, population)
		p0 = population
	}

	return years
}

/*
Buatlah program yang menentukan di tahun ke berapa sebuah kota akan mencapai atau
melebihi kapasitas penduduknya.Dengan ketentuan p0 adalah jumlah penduduk saat ini,
percent adalah persentase kenaikan jumlah penduduk pertahun,
aug adalah jumlah pendatang baru pertahun, dan p adalah kapasitas dari kota tersebut.
*/
func main() {
	p0 := 1000
	percent := 2
	aug := 50
	p := 1200

	years := nbYear(p0, percent, aug, p)
	fmt.Printf("It will need %d entire years.\n", years)
}
