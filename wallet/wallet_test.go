package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValue_ShouldReturnZeroOnEmptyWallet(t *testing.T) {
	w := NewWallet()
	assert.Equal(t, 0.0, w.GetValue(EUR, nil))
}

func TestGetValue_ShouldReturn10EurIfWalletContains10Eur(t *testing.T) {
	w := NewWallet()
	w.Add(EUR, 10)
	assert.Equal(t, 10.0, w.GetValue(EUR, nil))
}

func TestGetValue_ShouldReturn10UsdIfWalletContains10Usd(t *testing.T) {
	w := NewWallet()
	w.Add(USD, 10)
	assert.Equal(t, 10.0, w.GetValue(USD, nil))
}

type MockRate struct{}

func (r *MockRate) GetRate(fromCurrency currency, toCurrency currency) float64 {
	return 1.5
}

func TestGetValue_ShouldReturnTotalValueIfWalletContains10UsdAnd20Eur(t *testing.T) {
	w := NewWallet()
	w.Add(EUR, 20)
	w.Add(USD, 10)

	assert.Equal(t, 40.0, w.GetValue(USD, new(MockRate)))
}
