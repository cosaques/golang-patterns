package main

type Wallet struct {
	amounts map[string]float64
}

func (w *Wallet) GetValue(currency string) float64 {
	return w.amounts[currency]
}

func (w *Wallet) Add(currency string, amount float64) {
	w.amounts[currency] += amount
}

func NewWallet() *Wallet {
	return &Wallet{amounts: make(map[string]float64)}
}
