package day19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 1, PartOne(2))
	assert.Equal(t, 3, PartOne(3))
	assert.Equal(t, 1, PartOne(4))
	assert.Equal(t, 3, PartOne(5))
	assert.Equal(t, 5, PartOne(6))
	assert.Equal(t, 7, PartOne(7))
	assert.Equal(t, 1, PartOne(8))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 1830117, PartOne(3012210))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 2, PartTwo(5))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 1417887, PartTwo(3012210))
}
