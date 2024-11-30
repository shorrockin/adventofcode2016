package grid

type Coordinate struct {
	X int
	Y int
}

func At(x, y int) Coordinate {
	return Coordinate{x, y}
}

func (g Coordinate) North() Coordinate     { return Coordinate{g.X, g.Y - 1} }
func (g Coordinate) East() Coordinate      { return Coordinate{g.X + 1, g.Y} }
func (g Coordinate) NorthEast() Coordinate { return Coordinate{g.X + 1, g.Y - 1} }
func (g Coordinate) South() Coordinate     { return Coordinate{g.X, g.Y + 1} }
func (g Coordinate) SouthEast() Coordinate { return Coordinate{g.X + 1, g.Y + 1} }
func (g Coordinate) West() Coordinate      { return Coordinate{g.X - 1, g.Y} }
func (g Coordinate) SouthWest() Coordinate { return Coordinate{g.X - 1, g.Y + 1} }
func (g Coordinate) NorthWest() Coordinate { return Coordinate{g.X - 1, g.Y - 1} }

func (g Coordinate) Cardinals() []Coordinate {
	return []Coordinate{g.North(), g.East(), g.South(), g.West()}
}
func (g Coordinate) Intercardinals() []Coordinate {
	return []Coordinate{g.NorthEast(), g.SouthEast(), g.SouthWest(), g.NorthWest()}
}
func (g Coordinate) Compass() []Coordinate {
	return []Coordinate{g.North(), g.East(), g.NorthEast(), g.South(), g.SouthEast(), g.West(), g.SouthWest(), g.NorthWest()}
}
