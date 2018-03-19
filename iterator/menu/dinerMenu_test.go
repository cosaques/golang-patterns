package menu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDinerMenu_ShouldReturnNotEmtyMenu(t *testing.T) {
	assert.True(t, len(NewDinerMenu().GetItems()) > 0)
}

func TestDinerMenuIterator_ShouldReturnCorrectElements(t *testing.T) {
	list := []interface{}{1, 2}
	iterator := newDinerMenuIterator(list)

	assert.True(t, iterator.HasNext())
	assert.Equal(t, iterator.Next(), 1)

	assert.True(t, iterator.HasNext())
	assert.Equal(t, iterator.Next(), 2)

	assert.False(t, iterator.HasNext())
}
