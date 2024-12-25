package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	set := NewSet[string]()
	set.Add("Chris")
	set.Add("Shorrock")
	assert.True(t, set.Contains("Chris"))
	assert.True(t, set.Contains("Shorrock"))
	assert.False(t, set.Contains("John"))
	assert.Equal(t, 2, set.Size())

	set.Remove("Chris")
	assert.True(t, set.Contains("Shorrock"))
	assert.False(t, set.Contains("Chris"))
	assert.Equal(t, 1, set.Size())
}
