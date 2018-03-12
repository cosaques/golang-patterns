package wallet

type wallet struct {
	amounts map[currency]stock
}

func (w *wallet) GetValue(currency currency, rate Rate) stock {
	if rate == nil {
		return NewStock(w.amounts[currency].amount, currency)
	}
	return w.getTotalValue(currency, rate)
}

func (w *wallet) Add(amount stock) {
	w.amounts[amount.currency] = amount
}

func (w *wallet) getTotalValue(toCurrency currency, rate Rate) stock {
	total := NewStock(0, toCurrency)
	for currency, value := range w.amounts {
		tmpStock := NewStock(value.amount*computeRate(currency, toCurrency, rate), toCurrency)
		total, _ = AddStock(total, tmpStock)
	}
	return total
}

func computeRate(fromCurrency currency, toCurrency currency, rate Rate) float64 {
	if fromCurrency != toCurrency {
		return rate.GetRate(fromCurrency, toCurrency)
	}
	return 1.0
}

func NewWallet() *wallet {
	return &wallet{amounts: make(map[currency]stock)}
}
