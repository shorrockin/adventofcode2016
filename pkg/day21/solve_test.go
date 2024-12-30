package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, "decab", Solve("input.example.txt", "abcde", true))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, "gbhafcde", Solve("input.txt", "abcdefgh", true))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, "bcfaegdh", Solve("input.txt", "fbgdceah", false))
}
