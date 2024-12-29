package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 3, PartOne("input.example.txt"))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 22887907, PartOne("input.txt"))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, MAX_VALUE-7, PartTwo("input.example.txt"))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 109, PartTwo("input.txt"))
}
