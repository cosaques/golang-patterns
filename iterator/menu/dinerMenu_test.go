package menu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDinerMenu_ShouldReturnNotEmtyMenu(t *testing.T) {
	assert.True(t, len(NewDinerMenu().GetItems()) > 0)
}
