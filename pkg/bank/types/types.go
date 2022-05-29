package types

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type Currency string
type Money int

type Card struct {
	Balance    Money
	MinBalance Money
	Active     bool
}
