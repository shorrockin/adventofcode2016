package graph

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewGraph(t *testing.T) {
	g := NewGraph[string]()
	g.AddEdge("A", "B", 20)
	g.AddBidirectionalEdge("B", "C", 40)

	assert.Equal(t, 3, g.NodeCount())

	distance, exists := g.Distance("A", "B")
	assert.True(t, exists)
	assert.Equal(t, float64(20), distance)

	_, exists = g.Distance("B", "A")
	assert.False(t, exists)

	distance, exists = g.Distance("B", "C")
	assert.True(t, exists)
	assert.Equal(t, float64(40), distance)

	distance, exists = g.Distance("C", "B")
	assert.True(t, exists)
	assert.Equal(t, float64(40), distance)
}
