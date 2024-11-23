package graph

import (
	"adventofcode2016/pkg/utils"
	"math"
)

// given a graph and a source destination, returns a map
// of distances which indicate the shortest path to each
// possible node
func Dijkstra[T comparable](graph *Graph[T], source T) map[T]float64 {
	distances := make(map[T]float64)

	// initialize all distances to infinity
	for parent := range graph.edges {
		distances[parent] = math.Inf(1)

		// some children are not parent's so add them as well
		// to ensure we capture everything
		for child := range graph.edges[parent] {
			distances[child] = math.Inf(1)
		}
	}
	distances[source] = 0

	visited := utils.NewSet[T]()
	current, notDone := nextNode(distances, visited)

	for notDone {
		visited.Add(current)
		for neighbor, distance := range graph.edges[current] {
			alternative := distances[current] + distance
			if alternative < distances[neighbor] {
				distances[neighbor] = alternative
			}
		}

		current, notDone = nextNode(distances, visited)
	}

	return distances
}

func nextNode[T comparable](distances map[T]float64, visited utils.Set[T]) (T, bool) {
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
