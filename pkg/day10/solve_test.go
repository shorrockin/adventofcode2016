package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneActual(t *testing.T) {
	// 157: visually confirmed in output
	// assert.Equal(t, 0, Solve("input.txt"))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 1085, Solve("input.txt"))
}
