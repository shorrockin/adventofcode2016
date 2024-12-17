package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, "01100", Solve("10000", 20))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, "01110011101111011", Solve("11110010111001001", 272))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, "11001111011000111", Solve("11110010111001001", 35651584))
}
