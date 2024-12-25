package dijkstra

import (
	. "adventofcode2016/pkg/utils/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	g := NewGraph[string]()
	g.AddEdge("A", "B", 4)
	g.AddEdge("A", "C", 2)
	g.AddEdge("B", "C", 5)
	g.AddEdge("B", "D", 10)
	g.AddEdge("C", "E", 3)
	g.AddEdge("E", "D", 4)
	g.AddEdge("D", "F", 11)

	distances := Dijkstra(g, "A")
	assert.Equal(t, float64(4), distances["B"])
	assert.Equal(t, float64(2), distances["C"])
	assert.Equal(t, float64(9), distances["D"])
	assert.Equal(t, float64(5), distances["E"])
	assert.Equal(t, float64(20), distances["F"])
}
