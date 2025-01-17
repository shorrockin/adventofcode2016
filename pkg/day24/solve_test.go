package day24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 14, Solve("input.example.txt", true))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 460, Solve("input.txt", true))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 668, Solve("input.txt", false))
}
