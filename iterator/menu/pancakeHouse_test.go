package menu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPancakeHouse_ShouldNotBeEmpty(t *testing.T) {
	assert.NotNil(t, NewPancakeHouse().CreateIterator().Next())
}

func TestNewPancakeHouseIterator_ShouldReturnCorrectElements(t *testing.T) {
	m := map[string]interface{}{
		"a": 1,
		"b": 2,
	}
	iterator := newPancakeHouseIterator(m)

	assert.True(t, iterator.HasNext())
	assert.Equal(t, iterator.Next(), 1)

	assert.True(t, iterator.HasNext())
	assert.Equal(t, iterator.Next(), 2)

	assert.False(t, iterator.HasNext())
}
