package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 1985, PartOne("input.example.txt"))
}

func TestPartOne(t *testing.T) {
	assert.Equal(t, 48584, PartOne("input.txt"))
}

func TestPartTwo(t *testing.T) {
	assert.Equal(t, "563B6", PartTwo("input.txt"))
}
