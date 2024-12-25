package utils

import (
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
