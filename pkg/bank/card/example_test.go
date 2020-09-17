package card

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true},10_000_00)
	fmt.Println(result.Balance)

	//Output: 1000000
}

func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 10_000_00, Active: true},20_000_00)
	fmt.Println(result.Balance)

	//Output: 1000000
}

func ExampleWithdraw_inActive() {
	result := Withdraw(types.Card{Balance: 10_000_00, Active: false},500_00)
	fmt.Println(result.Balance)

	//Output: 1000000
}

func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 10_000_00, Active: false},25_000_00)
	fmt.Println(result.Balance)

	//Output: 1000000
}

func ExampleDeposit_positive(){
	card := types.Card{Balance: -100,Active:true}
	link := &card
	Deposit(link,100_000)
	fmt.Println(card.Balance)

	//Output : 99_900
}

func ExampleDeposit_inActive(){
	card := types.Card{Balance: 0,Active:false}
	link := &card
	Deposit(link,100_000)
	fmt.Println(card.Balance)

	//Output : 0
}

func ExampleDeposit_limit(){
	card := types.Card{Balance: 100,Active:true}
	link := &card
	Deposit(link,500_01_00)
	fmt.Println(card.Balance)

	//Output : 0
}
func ExampleAddBonus_positive(){
	card := types.Card{Balance: 100,MinBalance : 1_000_000 ,Active:true}
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)

	//Output : 246.00
}
func ExampleAddBonus_isActive(){
	card := types.Card{Balance: 100,MinBalance : 1_000_000 ,Active:false} 
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)

	//Output : 1000000
}

func ExampleAddBonus_isNegative(){
	card := types.Card{Balance: -100,MinBalance : 1_000_000 ,Active:true} 
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)

	//Output : 1000000
}

func ExampleAddBonus_limit(){
	card := types.Card{Balance: 100,MinBalance : 1000_000_000 ,Active:true} 
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)

	//Output : 1000000
}

func ExampleTotal(){
	cards := []types.Card{
		{
			Balance: 1_000, 
			Active: true,
		},
		{
			Balance: 2_000, 
			Active: true,
		},
		{
			Balance: 3_000, 
			Active: true,
		},
	}

	fmt.Print(Total(cards))

	//Output: 6000

}

func ExamplePaymentSources(){
	cards := []types.Card{
	{
		PAN:     "5058 xxxx xxxx 8888",
		Balance: 10_000_00,
		Active:  true,
	},
	{
		PAN:     "5058 xxxx xxxx 0000",
		Balance: -10_000_00,
		Active:  true,
	},
	{
		PAN:     "5058 xxxx xxxx 1111",
		Balance: 15_000_00,
		Active:  false,
	},
	}
	paymentSources := PaymentSources(cards)

	for _, v := range paymentSources {
		fmt.Println(v.Number)
	}

	//Output: 5058 xxxx xxxx 8888

}