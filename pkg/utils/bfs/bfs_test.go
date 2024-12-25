package bfs

import (
	"adventofcode2016/pkg/utils/graph"
	"adventofcode2016/pkg/utils/maps"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFS(t *testing.T) {
	g := graph.NewGraph[string]()
	g.AddEdge("A", "B", 1)
	g.AddEdge("A", "C", 1)
	g.AddEdge("B", "D", 1)
	g.AddEdge("C", "B", 1)
	g.AddEdge("D", "F", 1)
	g.AddEdge("E", "F", 1)

	neighbors := func(node string) []string {
		return maps.Keys(g.Edges[node])
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
