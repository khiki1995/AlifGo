package payment

import (
	"fmt"
	"bank/pkg/bank/types"
)
func ExampleMax(){
	payments := []types.Payment{
		{
			ID: 001, 
			Amount: 1000,
		},
		{
			ID: 002, 
			Amount: 5000,
		},
		{
			ID: 004, 
			Amount: 2000,
		},
	}

	fmt.Print(Max(payments))

	//Output: {2 5000}

}