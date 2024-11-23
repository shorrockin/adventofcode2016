package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.False(t, IsValidTriangle([]int{5, 10, 25}))
	assert.True(t, IsValidTriangle([]int{5, 5, 5}))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 869, PartOne("input.txt"))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 1544, PartTwo("input.txt"))
}
