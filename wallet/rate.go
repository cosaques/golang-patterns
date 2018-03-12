package wallet

type Rate interface {
	GetRate(fromCurrency currency, toCurrency currency) float64
}
