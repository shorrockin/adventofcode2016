package graph

import (
	"adventofcode2016/pkg/utils/collections"
)

type Graph[T comparable] struct {
	Edges map[T]map[T]float64
	Nodes collections.Set[T]
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{
		Edges: make(map[T]map[T]float64),
		Nodes: collections.NewSet[T](),
	}
}

func (g *Graph[T]) ConnectedTo(from T) []T {
	keys := make([]T, 0, len(g.Edges[from]))
	for key := range g.Edges[from] {
		keys = append(keys, key)
	}
	return keys
}

// add's a directed edge between two nodes in the graph
func (g *Graph[T]) AddEdge(from T, to T, weight float64) {
	if g.Edges[from] == nil {
		g.Edges[from] = make(map[T]float64)
	}
	g.Nodes.Add(from)
	g.Nodes.Add(to)
	g.Edges[from][to] = weight
}

// add's a bidirectional edge between two nodes in the graph
func (g *Graph[T]) AddBidirectionalEdge(from T, to T, weight float64) {
	g.AddEdge(from, to, weight)
	g.AddEdge(to, from, weight)
}

// true if there is a connection between the two edges
func (g Graph[T]) Exists(from T, to T) bool {
	value, exists := g.Edges[from]
	if !exists {
		return false
	}

	_, exists = value[to]
	if !exists {
		return false
	}

	return true
}

// returns the distance between the two edges, or -1, false if it
// doesn't exist
func (g Graph[T]) Distance(from, to T) (float64, bool) {
	if !g.Exists(from, to) {
		return -1, false
	}
	return g.Edges[from][to], true
}

func (g Graph[T]) Connected(from, to T) bool {
	_, connected := g.Distance(from, to)
	return connected
}

// returns the total number of unique nodes in this graph as
// the map can be hard to reason about if you add a single
// entry for A -> B this needs to return 2, even though the
// map size is 1
func (g Graph[T]) NodeCount() int {
	return len(g.Nodes)
}

func (g *Graph[T]) Values() []T {
	return g.Nodes.Values()
}
