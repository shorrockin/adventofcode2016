package day05

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func String[T []byte | [16]byte](values T) string {
	return fmt.Sprintf("%x", values)
}

func TestHashing(t *testing.T) {
	assert.Equal(t, "000008f82c5b3924a1ecbebf60344e00", String(Hash([]byte("abc"), 5017308)))
}

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, "18f47a30", PartOne([]byte("abc")))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, "f97c354d", PartOne([]byte("reyedfim")))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, "05ace8e3", PartTwo([]byte("abc")))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, "863dde27", PartTwo([]byte("reyedfim")))
}
