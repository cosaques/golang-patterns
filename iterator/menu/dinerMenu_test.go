package restaurant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDinerMenu_ShouldNotBeEmpty(t *testing.T) {
	assert.NotNil(t, NewDinerMenu().CreateIterator().Next())
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
