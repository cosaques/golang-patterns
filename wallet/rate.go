package wallet

type Rate interface {
	GetRate(fromCurrency string, toCurrency string) float64
}
