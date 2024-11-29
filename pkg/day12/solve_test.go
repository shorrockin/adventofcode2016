package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 42, Solve("input.example.txt", 0))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 318003, Solve("input.txt", 0))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 9227657, Solve("input.txt", 1))
}
