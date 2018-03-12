package wallet

type wallet struct {
	amounts map[string]float64
}

func (w *wallet) GetValue(currency string) float64 {
	return w.amounts[currency]
}

func (w *wallet) Add(currency string, amount float64) {
	w.amounts[currency] += amount
}

func NewWallet() *wallet {
	return &wallet{amounts: make(map[string]float64)}
}
