package day25

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 180, Solve("input.txt"))
}
