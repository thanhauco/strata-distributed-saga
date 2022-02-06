package types

import "fmt"

type Money struct {
	Amount   int64  `json:"amount_cents"`
	Currency string `json:"currency"`
}

func NewMoney(dollars float64, currency string) (Money, error) {
	if dollars < 0 {
		return Money{}, fmt.Errorf("money amount cannot be negative")
	}
	return Money{Amount: int64(dollars * 100), Currency: currency}, nil
}
