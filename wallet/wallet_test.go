package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValue_ShouldReturnZeroOnEmptyWallet(t *testing.T) {
	w := NewWallet()
	assert.Equal(t, NewStock(0, EUR), w.GetValue(EUR, nil))
}

func TestGetValue_ShouldReturn10EurIfWalletContains10Eur(t *testing.T) {
	w := NewWallet()
	w.Add(NewStock(10, EUR))
	assert.Equal(t, NewStock(10, EUR), w.GetValue(EUR, nil))
}

func TestGetValue_ShouldReturn10UsdIfWalletContains10Usd(t *testing.T) {
	w := NewWallet()
	w.Add(NewStock(10, USD))
	assert.Equal(t, NewStock(10, USD), w.GetValue(USD, nil))
}

type MockRate struct{}

func (r *MockRate) GetRate(fromCurrency currency, toCurrency currency) float64 {
	return 1.5
}

func TestGetValue_ShouldReturnTotalValueIfWalletContains10UsdAnd20Eur(t *testing.T) {
	w := NewWallet()
	w.Add(NewStock(20, EUR))
	w.Add(NewStock(10, USD))

	assert.Equal(t, NewStock(40, USD), w.GetValue(USD, new(MockRate)))
}
