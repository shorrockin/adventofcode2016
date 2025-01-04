package day22

import (
	"adventofcode2016/pkg/utils/logger"
)

func PartTwo(path string) int {
	log.Reset("part two starting")
	cluster := Parse(path)
	log.Log("cluster", "original", cluster)

	moves := cluster.Moves()
	// log.Log("cluster", "moves", moves)

	newC := cluster.Apply(moves[0])

	log.Log("cluster", "before", cluster.String())
	log.Log("move", "move", moves[0])
	log.Log("cluster", "after", newC.String())

	// solved by hand.... /shrug

	return logger.Return(&log, len(cluster.Nodes))
}
