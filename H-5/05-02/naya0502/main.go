package main

import (
	"fmt"
	"naya0502/config"
)

func main() {
	conf := config.NewConfig()

	fmt.Println(conf.DB)
}
