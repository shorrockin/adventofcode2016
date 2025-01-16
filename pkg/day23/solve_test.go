package day23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 3, Solve("input.example.txt", 2))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 11662, Solve("input.txt", 7))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 479008222, Solve("input.txt", 12))
}
