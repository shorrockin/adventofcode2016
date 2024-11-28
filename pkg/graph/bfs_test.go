package graph

import (
	"adventofcode2016/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFS(t *testing.T) {
	graph := NewGraph[string]()
	graph.AddEdge("A", "B", 1)
	graph.AddEdge("A", "C", 1)
	graph.AddEdge("B", "D", 1)
	graph.AddEdge("C", "B", 1)
	graph.AddEdge("D", "F", 1)
	graph.AddEdge("E", "F", 1)

	neighbors := func(node string) []string {
		return utils.Keys(graph.edges[node])
	}

	complete := func(node string) bool {
		return node == "F"
	}

	path, ok := BFS("A", neighbors, complete)
	assert.True(t, ok)
	assert.Equal(t, []string{"A", "B", "D", "F"}, path)

	path, ok = BFS("B", neighbors, complete)
	assert.True(t, ok)
	assert.Equal(t, []string{"B", "D", "F"}, path)
}
