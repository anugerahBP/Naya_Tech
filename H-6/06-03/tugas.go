package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Command struct {
	Nama string
	Usia int
}

func parser(input string) Command {
	data := strings.Split(input, ",")

	var nama string
	var usia int

	for _, item := range data {
		pair := strings.Split(item, ":")
		if len(pair) != 2 {
			continue
		}

		key := strings.TrimSpace(pair[0])
		value := strings.TrimSpace(pair[1])

		switch key {
		case "nama":
			nama = value
		case "usia":
			usia, _ = strconv.Atoi(value)
		}
	}

	return Command{
		Nama: nama,
		Usia: usia,
	}
}

func main() {
	input := "nama: joki siang sore, usia: 19, nama:bayu, usia:22"
	result := parser(input)

	fmt.Printf("Command{\n  Nama: \"%s\",\n  Usia: %d,\n}\n", result.Nama, result.Usia)
}
