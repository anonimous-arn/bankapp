package card

import (
	pay "bank/pkg/bank/payment"
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 1000000
}

func ExampleWithdraw_noMoney() {

	result := Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 0

}

func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 0, Active: false}, 10_000_00)
	fmt.Println(result.Active)
	// Output: false
}

func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 500_000, Active: true}, 30_000_00)
	fmt.Println(result.Balance)
	// Output: 500000
}

func ExampleDeposit_positive() {
	card := types.Card{MinBalance: 0, Active: true}
	Deposit(&card, 3000000)
	// Output: 3000000
}

func ExampleDeposit_inactive() {
	card := types.Card{Balance: 0, Active: false}
	Deposit(&card, 30_000_00)
	// Output: 0
}

func ExampleDeposit_limit() {
	card := types.Card{Balance: 10_000_00, Active: true}
	Deposit(&card, 800_000_00)
	// Output: 1000000
}

func ExampleAddBonus_inactive() {
	result := types.Card{MinBalance: 10_00_00, Active: true}
	AddBonus(&result, 3, 30, 365)

}

func ExampleTotal() {
	cards := []types.Card{

		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
	}
	fmt.Println(Total(cards))
	//Output: 3000000

}

func ExampleMax() {

	payments := []types.Payment{
		{
			ID:     12,
			Amount: 100,
		},
		{
			ID:     13,
			Amount: 300,
		},
		{
			ID:     14,
			Amount: 560,
		},
		{
			ID:     14,
			Amount: 5600,
		},
	}

	max := pay.Max(payments)
	fmt.Println(max)
	//Output: {14 5600}

}