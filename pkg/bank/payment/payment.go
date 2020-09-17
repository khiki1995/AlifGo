package payment

import "bank/pkg/bank/types"

func Max(payment []types.Payment) types.Payment {
	maxPayment := types.Money(0)
	index := 0
	for ind, payment := range payment{
		if(payment.Amount > maxPayment){
			maxPayment = payment.Amount
			index = ind
		}
	}
	return payment[index]
}