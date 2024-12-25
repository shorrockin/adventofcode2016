package grid

import (
	"adventofcode2016/pkg/utils/assert"
	"fmt"
	"math"
)

type Coord struct {
	X     int
	Y     int
	label string
}

func At(x, y int) Coord {
	return Coord{x, y, ""}
}

func (g Coord) WithLabel(label string) Coord {
	g.label = label
	return g
}

func (g Coord) North() Coord     { return At(g.X, g.Y-1) }
func (g Coord) East() Coord      { return At(g.X+1, g.Y) }
func (g Coord) NorthEast() Coord { return At(g.X+1, g.Y-1) }
func (g Coord) South() Coord     { return At(g.X, g.Y+1) }
func (g Coord) SouthEast() Coord { return At(g.X+1, g.Y+1) }
func (g Coord) West() Coord      { return At(g.X-1, g.Y) }
func (g Coord) SouthWest() Coord { return At(g.X-1, g.Y+1) }
func (g Coord) NorthWest() Coord { return At(g.X-1, g.Y-1) }

func (g Coord) Cardinals() []Coord {
	return []Coord{g.North(), g.East(), g.South(), g.West()}
}
func (g Coord) Intercardinals() []Coord {
	return []Coord{g.NorthEast(), g.SouthEast(), g.SouthWest(), g.NorthWest()}
}
func (g Coord) Compass() []Coord {
	return []Coord{g.North(), g.East(), g.NorthEast(), g.South(), g.SouthEast(), g.West(), g.SouthWest(), g.NorthWest()}
}

func (g Coord) Offset(offset Coord) Coord {
	return At(g.X+offset.X, g.Y+offset.Y)
}

func (g Coord) Move(offset Coord, times int) Coord {
	return At(g.X+offset.X*times, g.Y+offset.Y*times)
}

func (g Coord) Bounded(x, y int) Coord {
	if g.X >= x {
		g.X = g.X % x
	}
	if g.Y >= y {
		g.Y = g.Y % y
	}
	if g.X < 0 {
		g.X = x + g.X
	}
	if g.Y < 0 {
		g.Y = y + g.Y
	}
	return g
}

func (g Coord) TurnRight() Coord {
	switch g {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	default:
		panic(assert.Fail("couldn't turn right, needs to be a cardinal direction", "dir", g))
	}
}

func (g Coord) TurnLeft() Coord {
	switch g {
	case North:
		return West
	case East:
		return North
	case South:
		return East
	case West:
		return South
	default:
		panic(assert.Fail("couldn't turn left, needs to be a cardinal direction", "dir", g))
	}
}

func (g Coord) Distance(from Coord) int {
	return int(math.Abs(float64(g.X-from.X))) + int(math.Abs(float64(g.Y-from.Y)))
}

func (g Coord) CoordsInRange(maxRange int) []Coord {
	points := make([]Coord, 0, CoordsWithinRadius(maxRange))
	for dx := -maxRange; dx <= maxRange; dx++ {
		remaining := maxRange - int(math.Abs(float64(dx)))
		for dy := -remaining; dy <= remaining; dy++ {
			points = append(points, Coord{X: g.X + dx, Y: g.Y + dy})
		}
	}

	return points
}

func (g Coord) String() string {
	if g.label != "" {
		return fmt.Sprintf("%v(x:%d,y:%d)", g.label, g.X, g.Y)
	}
	return fmt.Sprintf("(x:%d,y:%d)", g.X, g.Y)
}

func CoordsWithinRadius(distance int) int {
	return 2*distance*distance + 2*distance + 1
}

var (
	North      = At(0, -1).WithLabel("North")
	East       = At(1, 0).WithLabel("East")
	South      = At(0, 1).WithLabel("South")
	West       = At(-1, 0).WithLabel("West")
	NorthWest  = At(-1, -1).WithLabel("NorthWest")
	NorthEast  = At(1, -1).WithLabel("NorthEast")
	SouthWest  = At(-1, 1).WithLabel("SouthWest")
	SouthEast  = At(1, 1).WithLabel("SouthEast")
	Directions = []Coord{North, East, South, West, NorthWest, NorthEast, SouthWest, SouthEast}
)
