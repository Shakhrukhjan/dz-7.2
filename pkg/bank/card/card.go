package card

import (
	"bank/pkg/bank/types"
)

const LimitBonus = 5_000

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {

	Bonus := (int(card.MinBalance) * percent * daysInMonth) / daysInYear
	if !card.Active {
		return
	}
	if card.Balance <= 0 {
		return
	}
	if Bonus > LimitBonus {
		Bonus = 0
	}
	card.Balance += types.Money(Bonus)
}
