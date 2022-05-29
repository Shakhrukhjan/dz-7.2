package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
)

func main() {

	template := &(types.Card{Balance: 20_000, MinBalance: 10_000, Active: true})
	card.AddBonus(template, 3, 30, 365)

}
