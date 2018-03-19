package menu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPancakeHouse_ShouldReturnNotEmptyMenu(t *testing.T) {
	assert.True(t, len(NewPancakeHouse().GetItems()) > 0)
}
