package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	assert.False(t, q.IsEmpty())

	value, ok := q.Pop()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
	assert.False(t, q.IsEmpty())

	value, ok = q.Pop()
	assert.Equal(t, 2, value)
	assert.True(t, ok)
	assert.False(t, q.IsEmpty())

	value, ok = q.Pop()
	assert.Equal(t, 3, value)
	assert.True(t, ok)
	assert.True(t, q.IsEmpty())

	_, ok = q.Pop()
	assert.False(t, ok)
}
