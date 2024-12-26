package astar

import (
	"adventofcode2016/pkg/utils/collections"
)

type Preference struct {
	allowBacktrack bool
	includeStart   bool
}

type Configurator func(*Preference)
type EndChecker[T comparable] func(node *collections.PqNode[T]) bool
type NeighborsRetriever[T comparable] func(node *collections.PqNode[T]) []T
type Heuristic[T comparable] func(node T, from *collections.PqNode[T]) float64

var ExcludeStart Configurator = func(p *Preference) {
	p.includeStart = false
}

var AllowBacktrack Configurator = func(p *Preference) {
	p.allowBacktrack = true
}

func AtEnd[T comparable](value T) EndChecker[T] {
	return func(node *collections.PqNode[T]) bool {
		return node.Contents == value
	}
}

func NoHeuristic[T comparable](node T, from *collections.PqNode[T]) float64 {
	return 0
}

func AStar[T comparable](start T, complete EndChecker[T], neighbors NeighborsRetriever[T], heuristic Heuristic[T], configs ...Configurator) []T {
	preferences := &Preference{
		allowBacktrack: false,
		includeStart:   true,
	}
	for _, configurator := range configs {
		configurator(preferences)
	}

	pq := collections.NewPriorityQueue[T]()
	visited := collections.NewSet[T]()
	pq.Push(start, 0, nil)

	for pq.Len() > 0 {
		current := pq.PopNode()

		if complete(current) {
			path := make([]T, 0)
			for current != nil {
				path = append([]T{current.Contents}, path...)
				current = current.Parent
			}

			if !preferences.includeStart && len(path) > 1 {
				return path[1:]
			} else {
				return path
			}
		}

		for _, neighbor := range neighbors(current) {
			if !preferences.allowBacktrack {
				if visited.Contains(neighbor) {
					continue
				}
				visited.Add(neighbor)
			}

			pq.Push(neighbor, heuristic(neighbor, current), current)
		}
	}

	// no path found, return empty array
	return make([]T, 0)
}
