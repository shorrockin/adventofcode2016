package day22

import (
	"adventofcode2016/pkg/utils/logger"
	"adventofcode2016/pkg/utils/slices"
)

func PartOne(path string) int {
	log.Reset("part one starting")
	cluster := Parse(path)
	combinations := slices.Combinations(cluster.Nodes[:], 2)
	count := slices.Reduce(combinations, 0, func(acc int, nodes []Node) int {
		if nodes[0].moveable(&nodes[1]) {
			acc++
		}
		if nodes[1].moveable(&nodes[0]) {
			acc++
		}
		return acc
	})

	return logger.Return(&log, count)
}
