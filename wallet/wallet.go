package wallet

type wallet struct {
	amounts map[string]float64
}

func (w *wallet) GetValue(currency string, rate Rate) float64 {
	result := 0.0
	if rate == nil {
		return w.amounts[currency]
	}
	for c, value := range w.amounts {
		r := 1.0
		if c != currency {
			r = rate.GetRate(c, currency)
		}
		result += value * r
	}
	return result
}

func (w *wallet) Add(currency string, amount float64) {
	w.amounts[currency] += amount
}

func NewWallet() *wallet {
	return &wallet{amounts: make(map[string]float64)}
}
