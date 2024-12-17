package utils

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInput(t *testing.T) {
	lines, err := ReadInput("utils.test.txt")
	assert.Nil(t, err)
	assert.Equal(t, len(lines), 2)
	assert.Equal(t, lines[0], "first line")
	assert.Equal(t, lines[1], "second line")
}

func TestReadInputError(t *testing.T) {
	lines, err := ReadInput("invalid.path.txt")
	assert.Nil(t, lines)
	assert.NotNil(t, err)
}

func TestKeys(t *testing.T) {
	data := map[string]int{
		"chris": 1,
		"john":  2,
		"jimbo": 3,
	}
	keys := Keys(data)
	assert.True(t, slices.Contains(keys, "chris"))
	assert.True(t, slices.Contains(keys, "john"))
	assert.True(t, slices.Contains(keys, "jimbo"))
	assert.Equal(t, 3, len(data))

}
