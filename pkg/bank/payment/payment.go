package payment

import "bank/pkg/bank/types"


func Max(payments []types.Payment) types.Payment {
	max := payments[0]
	for i:=0;i<len(payments);i++ {
		if payments[i].Amount>max.Amount{
			max = payments[i]
		}
	}
	return max
}

func Save(payments []types.Payment) []types.Payment {
	var arrPayments []types.Payment
	for i:=0; i<len(payments);i++ {
		arrPayments = append(arrPayments, payments[i])
	}
	return arrPayments
}
