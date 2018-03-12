package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValue_ShouldReturnZeroOnEmptyWallet(t *testing.T) {
	w := NewWallet()
	assert.Equal(t, 0.0, w.GetValue("EUR"))
}

func TestGetValue_ShouldReturn10EurIfWalletContains10Eur(t *testing.T) {
	w := NewWallet()
	w.Add("EUR", 10)
	assert.Equal(t, 10.0, w.GetValue("EUR"))
}

func TestGetValue_ShouldReturn10UsdIfWalletContains10Usd(t *testing.T) {
	w := NewWallet()
	w.Add("USD", 10)
	assert.Equal(t, 10.0, w.GetValue("USD"))
}
