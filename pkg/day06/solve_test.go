package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExampleInput(t *testing.T) {
	assert.Equal(t, "easter", Solve("input.example.txt", true))
}

func TestPartTwoExampleInput(t *testing.T) {
	assert.Equal(t, "advent", Solve("input.example.txt", false))
}

func TestPartOneActualInput(t *testing.T) {
	assert.Equal(t, "umejzgdw", Solve("input.txt", true))
}

func TestPartTwoActualInput(t *testing.T) {
	assert.Equal(t, "aovueakv", Solve("input.txt", false))
}
