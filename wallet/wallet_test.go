package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Should(t *testing.T) {
	w := &Wallet{}
	assert.Equal(t, 0.0, w.GetValue("EUR"))
}
