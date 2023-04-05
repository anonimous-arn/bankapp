package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

// Функция зачисления средств
func Deposit(card *types.Card, amount types.Money) {

	if amount > 500_000_00 || card.Active == false {
		fmt.Println(card.Balance)
		return
	} else {
		card.Balance += amount
		fmt.Println(card.Balance)
	}

}

// Зачисление бонуса на минималӣнҷй остаток
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {

	if card.MinBalance >= 10_000 {
		result :=int(card.MinBalance) + (int(card.MinBalance) * percent / 100 * daysInMonth / daysInYear)
		fmt.Println(result)
	} else {
		fmt.Println( "Остаток суммы составило на балансе клиента менӣше 10000 сомони на этот месяц")
	} 
}

// Данная функқия посчитает баланс всех активнҷх карт
func Total(cards []types.Card) types.Money {
var sum types.Money

for _, cards := range cards {
	sum += cards.Balance
}

	return sum
}


