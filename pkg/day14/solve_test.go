package day14

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasher(t *testing.T) {
	hasher := NewHasher("abc", false)
	assert.Equal(t, 12, strings.Index(hasher.At(18), "888"))
	assert.Equal(t, 10, strings.Index(hasher.At(39), "eee"))
	assert.Equal(t, 2, strings.Index(hasher.At(816), "eeeee"))
	assert.Equal(t, 26, strings.Index(hasher.At(92), "999"))
	assert.Equal(t, 9, strings.Index(hasher.At(200), "99999"))
}

func TestTripleAndQuintuple(t *testing.T) {
	assert.True(t, HasQuintuple("cc3888887a5e4e", '8'))
	assert.False(t, HasQuintuple("cc3888867a5e4e", '8'))

	value, ok := HasTriple("cc38887a5e4e")
	assert.True(t, ok)
	assert.Equal(t, 3, value)
}

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 22728, Solve("abc", false))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 22551, Solve("abc", true))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 35186, Solve("jlmsuwbz", false))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 22429, Solve("jlmsuwbz", true))
}
