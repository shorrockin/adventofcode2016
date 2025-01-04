package day22

import (
	"adventofcode2016/pkg/utils/assert"
	"adventofcode2016/pkg/utils/colors"
	"fmt"
	"strings"
)

type Move [2]int
type Node [4]int

func (n *Node) x() int                 { return n[0] }
func (n *Node) y() int                 { return n[1] }
func (n *Node) Size() int              { return n[2] }
func (n *Node) Used() int              { return n[3] }
func (n *Node) Free() int              { return n.Size() - n.Used() }
func (n *Node) moveable(to *Node) bool { return n.Used() != 0 && n.Used() <= to.Free() }

type Cluster struct {
	width  int
	height int
	Nodes  []Node
	target [2]int
}

func (c *Cluster) index(x, y int) int { return x + (c.width * y) }

func (c Cluster) Apply(move Move) Cluster {
	nodes := make([]Node, len(c.Nodes))
	copy(nodes, c.Nodes)
	c.Nodes = nodes

	from := &c.Nodes[move[0]]
	to := &c.Nodes[move[1]]
	assert.True(from.moveable(to), "from is not moveable to", "from", from, "to", to)

	to[3] += from[3]
	from[3] = 0

	// c.history = append(c.history, move)

	return c
}

func (c *Cluster) Moves() []Move {
	out := make([]Move, 0)
	for _, node := range c.Nodes {
		out = c.appendMove(&node, 0, -1, out)
		out = c.appendMove(&node, 0, 1, out)
		out = c.appendMove(&node, 1, 0, out)
		out = c.appendMove(&node, -1, 0, out)
	}
	return out
}

func (c *Cluster) String() string {
	var out strings.Builder

	out.WriteString("\n")
	for row := range c.height {
		for col := range c.width {
			idx := c.index(col, row)
			node := c.Nodes[idx]
			used := fmt.Sprintf("%3d", node.Used())
			if node.Used() == 0 {
				used = colors.Green(used)
			}
			free := fmt.Sprintf("%-3d", node.Free())
			if node.Used() == 0 {
				free = colors.Yellow(free)
			}
			out.WriteString(fmt.Sprintf("%s/%s", used, free))
		}
		out.WriteString("\n")
	}

	return out.String()
}

func (c *Cluster) appendMove(to *Node, offsetX, offsetY int, moves []Move) []Move {
	if from := c.index(to.x()+offsetX, to.y()+offsetY); from >= 0 && from < len(c.Nodes) {
		if c.Nodes[from].moveable(to) {
			return append(moves, Move{from, c.index(to.x(), to.y())})
		}
	}
	return moves
}
