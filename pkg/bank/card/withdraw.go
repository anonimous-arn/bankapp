package card

import (
	"bank/pkg/bank/types"
)

func Withdraw(card types.Card, amount types.Money) types.Card {
	if card.Balance == 0 || card.Balance < amount || card.Active == false || amount > 20_000_00 {

		return card

	}
	card.Balance = card.Balance - amount
	return card

}

