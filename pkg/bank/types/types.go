package types

// The sum of money in cent, kopeyka, diram
type Money int64

// The code of currency
type Currency string 

// The currency code

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"

)


// The PAN of cards
type PAN string

type Card struct {
	ID  int
	PAN PAN
	Balance Money
	MinBalance Money
	Currency Currency
	Color  string
	Name string
	Active bool
}

type Payment struct {

	ID  int
	Amount Money

}