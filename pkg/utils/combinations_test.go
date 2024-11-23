package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {
	combos := Combinations([]string{"A", "B", "C", "D"}, 2)
	assert.Equal(t, 6, len(combos))
	assert.Equal(t, []string{"A", "B"}, combos[0])
	assert.Equal(t, []string{"A", "C"}, combos[1])
	assert.Equal(t, []string{"A", "D"}, combos[2])
	assert.Equal(t, []string{"B", "C"}, combos[3])
	assert.Equal(t, []string{"B", "D"}, combos[4])
	assert.Equal(t, []string{"C", "D"}, combos[5])
}
