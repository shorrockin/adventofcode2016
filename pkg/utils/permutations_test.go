package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	input := []int{1, 2, 3}
	permutations := Permutations(input)
	assert.Equal(t, 6, len(permutations))
	assert.Equal(t, []int{1, 2, 3}, permutations[0])
	assert.Equal(t, []int{3, 1, 2}, permutations[5])
}

func TestPermutationsFiltered(t *testing.T) {
	input := []int{1, 2, 3}
	permutations := PermutationsFiltered(input, func(values []int, at int) bool {
		// skip permutations with 2 in the second position
		if at >= 1 {
			return values[1] != 2
		}
		return true
	})
	// should be the same as above, but [1, 2, 3] should be
	// removed as should [3, 2, 1]
	assert.Equal(t, 4, len(permutations))
	assert.Equal(t, []int{1, 3, 2}, permutations[0])
	assert.Equal(t, []int{3, 1, 2}, permutations[3])
}
