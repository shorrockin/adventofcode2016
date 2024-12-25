package tsp

import (
	"adventofcode2016/pkg/utils/assert"
	"adventofcode2016/pkg/utils/collections"
	. "adventofcode2016/pkg/utils/graph"
	"adventofcode2016/pkg/utils/maps"
	"adventofcode2016/pkg/utils/slices"
	"math"
)

type TSPPath[T comparable] struct {
	nodes     []T
	distances map[T]float64
	distance  float64
	visited   collections.Set[T]
}

func NewPath[T comparable](shortest bool) *TSPPath[T] {
	var distance float64
	switch shortest {
	case true:
		distance = math.Inf(1)
	case false:
		distance = math.Inf(-1)
	}

	return &TSPPath[T]{
		nodes:     make([]T, 0),
		distances: make(map[T]float64),
		distance:  distance,
		visited:   collections.NewSet[T](),
	}
}

func (p *TSPPath[T]) Copy() *TSPPath[T] {
	return &TSPPath[T]{
		nodes:     slices.Copy(p.nodes),
		distances: maps.Copy(p.distances),
		distance:  p.distance,
		visited:   p.visited.Copy(),
	}
}

func (p *TSPPath[T]) Distance() float64 {
	return p.distance
}

func (p *TSPPath[T]) Last() T {
	assert.Assert(len(p.nodes) != 0, "unable to get last node, length 0")
	return p.nodes[len(p.nodes)-1]
}

func (p *TSPPath[T]) Size() int {
	return len(p.nodes)
}

func (p *TSPPath[T]) BetterThan(other *TSPPath[T], shortest bool) bool {
	switch shortest {
	case true:
		return p.distance < other.distance
	default:
		return p.distance > other.distance
	}
}

func (p *TSPPath[T]) Push(node T, distance float64) {
	assert.Refute(p.visited.Contains(node), "can not add a node more than once", node)
	if len(p.visited) == 0 {
		p.distance = 0
	}

	p.visited.Add(node)
	p.distances[node] = distance
	p.distance += distance
	p.nodes = append(p.nodes, node)
}

func (p *TSPPath[T]) Pop() {
	assert.Assert(p.visited.Size() > 1, "you have no nodes to pop")
	removing := p.Last()

	assert.NotNil(p.distances[removing], "element being removed not in distances")
	distance := p.distances[removing]

	p.visited.Remove(removing)
	delete(p.distances, removing)
	p.distance -= distance
	p.nodes = p.nodes[:len(p.nodes)-1]
}

func (p *TSPPath[T]) Contains(node T) bool {
	return p.visited.Contains(node)
}

// implementation of a traveling salesman problem for a graph
// which calculates the best path to travel for all nodes in
// a graph.
func TSP[T comparable](graph *Graph[T], shortest bool) *TSPPath[T] {
	best := NewPath[T](shortest)

	for node := range graph.Edges {
		working := NewPath[T](shortest)
		working.Push(node, 0)
		best = recurseTSP(graph, best, working, shortest)
	}

	return best
}

func recurseTSP[T comparable](
	graph *Graph[T],
	best *TSPPath[T],
	working *TSPPath[T],
	shortest bool) *TSPPath[T] {

	subGraph := graph.Edges[working.Last()]
	for next := range subGraph {
		// we can not visit a location twice, skip if we've
		// already been here
		if working.Contains(next) {
			continue
		}

		working.Push(next, subGraph[next])

		bestBetter := best.BetterThan(working, shortest)
		fullGraph := (working.Size() == graph.NodeCount())

		// only continue to work on our working option if it's still
		// better than our best option, we can short circuit if at
		// any point this partial path becomes worse, only works
		// when iterating the shorter path
		if shortest && bestBetter {
			working.Pop()
			continue
		}

		if !fullGraph {
			best = recurseTSP(graph, best, working, shortest)
		} else if !bestBetter {
			best = working.Copy()
		}

		working.Pop()
	}

	return best
}
