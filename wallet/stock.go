package wallet

import (
	"errors"
)

type stock struct {
	amount   float64
	currency currency
}

func AddStock(s1 stock, s2 stock) (stock stock, err error) {
	if s1.currency == s2.currency {
		stock = NewStock(s1.amount+s2.amount, s1.currency)
		return
	}

	err = CurrencyMismatchError
	return
}

func NewStock(amount float64, currency currency) stock {
	return stock{amount, currency}
}

var CurrencyMismatchError = errors.New("Currency mismatch")
