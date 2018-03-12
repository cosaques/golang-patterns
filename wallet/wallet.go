package wallet

type wallet struct {
	amounts map[currency]float64
}

func (w *wallet) GetValue(currency currency, rate Rate) float64 {
	if rate == nil {
		return w.amounts[currency]
	}
	return w.getTotalValue(currency, rate)
}

func (w *wallet) Add(currency currency, amount float64) {
	w.amounts[currency] += amount
}

func (w *wallet) getTotalValue(toCurrency currency, rate Rate) float64 {
	total := 0.0
	for currency, value := range w.amounts {
		total += value * computeRate(currency, toCurrency, rate)
	}
	return total
}

// ternaire
func computeRate(fromCurrency currency, toCurrency currency, rate Rate) float64 {
	if fromCurrency != toCurrency {
		return rate.GetRate(fromCurrency, toCurrency)
	}
	return 1.0
}

func NewWallet() *wallet {
	return &wallet{amounts: make(map[currency]float64)}
}
