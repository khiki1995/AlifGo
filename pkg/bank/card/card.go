package card

import "bank/pkg/bank/types"

func IssueCard(currency types.Currency, color string, name string) types.Card {
	return types.Card {
		ID : 1000,
		PAN : "5058 xxxx xxxx 0001",
		Balance : 0,
		MinBalance : 0,
		Currency : currency,
		Color : color,
		Name : name,
		Active : true,
	}
}

func Withdraw(card types.Card, amount types.Money) types.Card {
	if (card.Active && card.Balance >= amount && amount > 0 && amount < 20_000_00){
		card.Balance = card.Balance - amount
	}
	return card
}

func Deposit(card *types.Card, amount types.Money) {
	if (!card.Active){
		return 
	}
	if (amount < 0 || amount > 500_00_00){
		return 
	}
	card.Balance += amount
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if (!card.Active){
		return 
	}
	if (card.Balance <= 0){
		return 
	}
	bonus := int(card.MinBalance)*percent/100.0*daysInMonth/daysInYear
	if bonus > 500_000 {
		return
	}
	card.Balance += types.Money(bonus) 
}

func Total(cards []types.Card) types.Money{
	totalSum := types.Money(0)
	for _, value := range cards {
		totalSum += value.Balance
	}
	return totalSum
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var PaymentSources []types.PaymentSource
	for _,card := range cards {
		if card.Balance > 0 && card.Active {
			PaymentSources = append(PaymentSources,types.PaymentSource{Number : string(card.PAN) ,Type: "card" ,Balance : card.Balance})
		}
	}
	return PaymentSources
}