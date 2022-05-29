package card

import (
	"bank/pkg/bank/types"
)

func ExampleAddBonus_norm() {
	template := &(types.Card{MinBalance: 10_000, Active: true})
	AddBonus(template, 3, 30, 365)
	// Output: 2465
}
func ExampleAddBonus_noActive() {
	template := &(types.Card{MinBalance: 10_000, Active: false})
	AddBonus(template, 3, 30, 365)
	// Output: 10000
}
func ExampleAddBonus_negative() {
	template := &(types.Card{MinBalance: -10_000, Active: true})
	AddBonus(template, 3, 30, 365)
	// Output: 10000
}
func ExampleAddBonus_limit() {
	template := &(types.Card{MinBalance: 10_000, Active: true})
	AddBonus(template, 3, 30, 365)
	// Output: 0
}
