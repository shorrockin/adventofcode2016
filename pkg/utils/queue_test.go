package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.False(t, q.IsEmpty())

	value, ok := q.Dequeue()
	assert.Equal(t, 1, value)
	assert.True(t, ok)
	assert.False(t, q.IsEmpty())

	value, ok = q.Dequeue()
	assert.Equal(t, 2, value)
	assert.True(t, ok)
	assert.False(t, q.IsEmpty())

	value, ok = q.Dequeue()
	assert.Equal(t, 3, value)
	assert.True(t, ok)
	assert.True(t, q.IsEmpty())

	_, ok = q.Dequeue()
	assert.False(t, ok)
}
