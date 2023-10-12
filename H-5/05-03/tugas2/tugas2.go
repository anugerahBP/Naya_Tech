package main

import (
	"fmt"
	"tugas2/fighters"
)

func main() {
	f1 := fighters.Fighter{Name: "Lew", Health: 10, DamagePerAttack: 2}
	f2 := fighters.Fighter{Name: "Harry", Health: 5, DamagePerAttack: 4}
	winner := fighters.DeclareWinner(f1, f2, "Lew")

	if winner != "" {
		fmt.Printf("%s wins.\n", winner)
	} else {
		fmt.Println("The battle ended in a draw.")
	}
}
