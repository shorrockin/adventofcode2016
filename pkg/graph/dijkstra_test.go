package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	graph := NewGraph[string]()
	graph.AddEdge("A", "B", 4)
	graph.AddEdge("A", "C", 2)
	graph.AddEdge("B", "C", 5)
	graph.AddEdge("B", "D", 10)
	graph.AddEdge("C", "E", 3)
	graph.AddEdge("E", "D", 4)
	graph.AddEdge("D", "F", 11)

	distances := Dijkstra(graph, "A")
	assert.Equal(t, float64(4), distances["B"])
	assert.Equal(t, float64(2), distances["C"])
	assert.Equal(t, float64(9), distances["D"])
	assert.Equal(t, float64(5), distances["E"])
	assert.Equal(t, float64(20), distances["F"])
}
