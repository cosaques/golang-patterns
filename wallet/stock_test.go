package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd_ShouldAddSameCurrencies(t *testing.T) {
	stock, _ := AddStock(NewStock(5, EUR), NewStock(15, EUR))
	assert.Equal(t, NewStock(20, EUR), stock)
}
