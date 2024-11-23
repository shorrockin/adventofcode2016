package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTSPPath(t *testing.T) {
	path := NewPath[string](true)
	assert.Equal(t, 0, 0, path.distance)

	path.Push("A", 1)
	path.Push("B", 2)
	path.Push("C", 3)

	assert.Equal(t, 3, path.Size())
	assert.Equal(t, float64(6), path.distance)
	assert.True(t, path.Contains("C"))

	path.Pop()

	assert.Equal(t, 2, path.Size())
	assert.Equal(t, float64(3), path.distance)
	assert.False(t, path.Contains("C"))
}

func TestTSP(t *testing.T) {
	graph := NewGraph[string]()
	graph.AddBidirectionalEdge("London", "Belfast", 518)
	graph.AddBidirectionalEdge("London", "Dublin", 464)
	graph.AddBidirectionalEdge("Dublin", "Belfast", 141)

	path := TSP(graph, true)
	assert.Equal(t, 3, len(path.nodes))
	assert.Equal(t, float64(605), path.distance)
	assert.Equal(t, "Dublin", path.nodes[1])

	path = TSP(graph, false)
	assert.Equal(t, 3, len(path.nodes))
	assert.Equal(t, float64(982), path.distance)
	assert.Equal(t, "London", path.nodes[1])
}
