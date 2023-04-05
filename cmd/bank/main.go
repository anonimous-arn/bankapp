package main

import (
	t "bank/pkg/bank/card"
	"bank/pkg/bank/types"
)

func main() {
	card := types.Card{ /*Balance: -10_000_00*/ MinBalance: 9_000, Active: false}
	//result := card.Withdraw(types.Card{Balance: 20_000_00, Active: false}, 30_000_00)
	//t.Deposit(&card,500_000_00)
	t.AddBonus(&card, 3, 30, 365)
}
