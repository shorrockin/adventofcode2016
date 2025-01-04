package day22

import (
	"adventofcode2016/pkg/utils"
	"regexp"
)

func Parse(path string) Cluster {
	nodes := make([]Node, 0)
	pattern := regexp.MustCompile(`\/dev\/grid\/node\-x(\d+)\-y(\d+)\s+(\d+)T\s+(\d+)T.*`)
	width := 0
	height := 0

	for _, line := range utils.MustReadInput(path) {
		if matches := pattern.FindStringSubmatch(line); matches != nil {
			x, y, size, used := utils.MustAtoi(matches[1]), utils.MustAtoi(matches[2]), utils.MustAtoi(matches[3]), utils.MustAtoi(matches[4])
			nodes = append(nodes, Node{x, y, size, used})
			width = utils.Max(width, x)
			height = utils.Max(height, y)
		}
	}

	cluster := Cluster{width + 1, height + 1, make([]Node, (width+1)*(height+1)), [2]int{width, 0}}
	// log.Log("making cluster", "width", width, "height", height, "len", len(cluster.nodes))

	for _, node := range nodes {
		// log.Log("inserting", "x", node.x(), "y", node.y(), logger.IndentOnce)
		cluster.Nodes[cluster.index(node.x(), node.y())] = node
	}

	return cluster
}
