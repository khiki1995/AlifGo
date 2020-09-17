package main
import (
	"fmt"
	"bank/pkg/bank/types"
)
func main() {
	newCard := types.Card{Balance: 100,MinBalance : 1_000_000 ,Active:true}
	AddBonus(&newCard,3,30,365)
	fmt.Printf("%v",newCard.Balance)
}
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if (!card.Active){
		return 
	}
	if (card.Balance <= 0){
		return 
	}
	bonus := card.MinBalance*types.Money(percent)/100.0*types.Money(daysInMonth)/types.Money(daysInYear)
	if bonus > 500_000 {
		return
	}
	card.Balance += bonus 
}