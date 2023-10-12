package fighters

import "fmt"

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func DeclareWinner(f1 Fighter, f2 Fighter, firstAttacker string) string {
	turn := firstAttacker

	for f1.Health > 0 && f2.Health > 0 {
		if turn == f1.Name {
			f2.Health -= f1.DamagePerAttack
			fmt.Printf("%s attacks %s; %s now has %d health.\n", f1.Name, f2.Name, f2.Name, f2.Health)
		} else {
			f1.Health -= f2.DamagePerAttack
			fmt.Printf("%s attacks %s; %s now has %d health.\n", f2.Name, f1.Name, f1.Name, f1.Health)
		}

		if f1.Health <= 0 {
			return f2.Name
		} else if f2.Health <= 0 {
			return f1.Name
		}

		turn = toggleTurn(turn)
	}

	return ""
}

func toggleTurn(turn string) string {
	if turn == "Lew" {
		return "Harry"
	}
	return "Lew"
}
