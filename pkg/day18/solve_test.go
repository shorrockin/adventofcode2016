package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, ".^^^^", next("..^^."))
	assert.Equal(t, "..^^...^^^", next("^.^^.^.^^."))
	assert.Equal(t, 38, Solve(".^^.^.^^^^", 9))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 1913, Solve("^.^^^..^^...^.^..^^^^^.....^...^^^..^^^^.^^.^^^^^^^^.^^.^^^^...^^...^^^^.^.^..^^..^..^.^^.^.^.......", 39))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 19993564, Solve("^.^^^..^^...^.^..^^^^^.....^...^^^..^^^^.^^.^^^^^^^^.^^.^^^^...^^...^^^^.^.^..^^..^..^.^^.^.^.......", 399999))
}
