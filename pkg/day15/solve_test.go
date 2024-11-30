package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 5, Solve("input.example.txt"))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 317371, Solve("input.partone.txt"))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 2080951, Solve("input.parttwo.txt"))
}
