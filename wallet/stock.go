package wallet

type stock struct {
	amount   float64
	currency currency
}

func AddStock(s1 stock, s2 stock) (stock stock, err error) {
	if s1.currency == s2.currency {
		stock = NewStock(20, EUR)
		return
	}

	return
}

func NewStock(amount float64, currency currency) stock {
	return stock{amount, currency}
}
