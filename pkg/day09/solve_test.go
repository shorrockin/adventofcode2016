package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecompressPartOneExample(t *testing.T) {
	assert.Equal(t, 6, Decompress("ADVENT", PartOne))
	assert.Equal(t, 7, Decompress("A(1x5)BC", PartOne))
	assert.Equal(t, 9, Decompress("(3x3)XYZ", PartOne))
	assert.Equal(t, 11, Decompress("A(2x2)BCD(2x2)EFG", PartOne))
}

func TestDecompressPartTwoExample(t *testing.T) {
	assert.Equal(t, len("XYZXYZXYZ"), Decompress("(3x3)XYZ", PartTwo))
	assert.Equal(t, len("XABCABCABCABCABCABCY"), Decompress("X(8x2)(3x3)ABCY", PartTwo))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 138735, Solve("input.txt", PartOne))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 11125026826, Solve("input.txt", PartTwo))
}
