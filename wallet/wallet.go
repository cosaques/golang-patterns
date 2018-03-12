package wallet

type wallet struct {
	amounts map[string]float64
}

func (w *wallet) GetValue(currency string, rate Rate) float64 {
	if rate == nil {
		return w.amounts[currency]
	}
	return w.getTotalValue(currency, rate)
}

func (w *wallet) Add(currency string, amount float64) {
	w.amounts[currency] += amount
}

func (w *wallet) getTotalValue(toCurrency string, rate Rate) float64 {
	total := 0.0
	for currency, value := range w.amounts {
		total += value * computeRate(currency, toCurrency, rate)
	}
	return total
}

// ternaire
func computeRate(fromCurrency string, toCurrency string, rate Rate) float64 {
	if fromCurrency != toCurrency {
		return rate.GetRate(fromCurrency, toCurrency)
	}
	return 1.0
}

func NewWallet() *wallet {
	return &wallet{amounts: make(map[string]float64)}
}
