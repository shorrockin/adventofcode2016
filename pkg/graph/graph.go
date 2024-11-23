package graph

import "adventofcode2016/pkg/utils"

type Graph[T comparable] struct {
	edges map[T]map[T]float64
	nodes utils.Set[T]
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{
		edges: make(map[T]map[T]float64),
		nodes: utils.NewSet[T](),
	}
}

// add's a directed edge between two nodes in the graph
func (g *Graph[T]) AddEdge(from T, to T, weight float64) {
	if g.edges[from] == nil {
		g.edges[from] = make(map[T]float64)
	}
	g.nodes.Add(from)
	g.nodes.Add(to)
	g.edges[from][to] = weight
}

// add's a bidirectional edge between two nodes in the graph
func (g *Graph[T]) AddBidirectionalEdge(from T, to T, weight float64) {
	g.AddEdge(from, to, weight)
	g.AddEdge(to, from, weight)
}

// true if there is a connection between the two edges
func (g Graph[T]) Exists(from T, to T) bool {
	value, exists := g.edges[from]
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
func (g Graph[T]) Distance(from T, to T) (float64, bool) {
	if !g.Exists(from, to) {
		return -1, false
	}
	return g.edges[from][to], true
}

// returns the total number of unique nodes in this graph as
// the map can be hard to reason about if you add a single
// entry for A -> B this needs to return 2, even though the
// map size is 1
func (g Graph[T]) NodeCount() int {
	return len(g.nodes)
}
