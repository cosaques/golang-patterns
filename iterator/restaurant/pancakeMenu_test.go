package restaurant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPancakeMenu_ShouldNotBeEmpty(t *testing.T) {
	assert.NotNil(t, NewPancakeMenu().CreateIterator().Next())
}

func TestNewPancakeMenuIterator_ShouldReturnCorrectElements(t *testing.T) {
	m := map[string]interface{}{
		"a": 1,
		"b": 2,
	}
	iterator := newPancakeMenuIterator(m)

	assert.True(t, iterator.HasNext())
	assert.Equal(t, iterator.Next(), 1)

	assert.True(t, iterator.HasNext())
	assert.Equal(t, iterator.Next(), 2)

	assert.False(t, iterator.HasNext())
}
