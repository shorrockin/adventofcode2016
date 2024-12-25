package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriorityQueueFunctions(t *testing.T) {
	pq := NewPriorityQueue[string]()

	pq.Push("2, 0", 2, nil)
	pq.Push("3, 0", 3, nil)
	pq.Push("1, 0", 1, nil)
	pq.Push("0, 0", 0, nil)

	assert.Equal(t, 4, pq.Len())
	assert.Equal(t, "0, 0", pq.Pop())
	assert.Equal(t, "1, 0", pq.Pop())
	assert.Equal(t, "2, 0", pq.Pop())
	assert.Equal(t, "3, 0", pq.Pop())
}
