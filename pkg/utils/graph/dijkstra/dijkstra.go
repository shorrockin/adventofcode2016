package dijkstra

import (
	"adventofcode2016/pkg/utils/collections"
	. "adventofcode2016/pkg/utils/graph"
	"math"
)

// given a graph and a source destination, returns a map
// of distances which indicate the shortest path to each
// possible node
func Dijkstra[T comparable](graph *Graph[T], source T) map[T]float64 {
	distances := make(map[T]float64, len(graph.Nodes))

	// initialize all distances to infinity
	for parent := range graph.Edges {
		distances[parent] = math.Inf(1)

		// some children are not parent's so add them as well
		// to ensure we capture everything
		for child := range graph.Edges[parent] {
			distances[child] = math.Inf(1)
		}
	}
	distances[source] = 0

	visited := collections.NewSet[T]()
	current, notDone := nextNode(distances, visited)

	for notDone {
		visited.Add(current)
		for neighbor, distance := range graph.Edges[current] {
			alternative := distances[current] + distance
			if alternative < distances[neighbor] {
				distances[neighbor] = alternative
			}
		}

		current, notDone = nextNode(distances, visited)
	}

	return distances
}

func nextNode[T comparable](distances map[T]float64, visited collections.Set[T]) (T, bool) {
	var currentNext T
	currentDistance := math.Inf(1)

	for node, distance := range distances {
		if distance < currentDistance && !visited.Contains(node) {
			currentNext = node
			currentDistance = distance
		}
	}

	if currentDistance == math.Inf(1) {
		return currentNext, false
	}

	return currentNext, true
}
