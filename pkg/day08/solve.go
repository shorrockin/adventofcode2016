package day08

import (
	"adventofcode2016/pkg/assert"
	"adventofcode2016/pkg/grid"
	"adventofcode2016/pkg/utils"
	"fmt"
	"strings"
)

type Screen []grid.Coordinate
type Command func(Screen) Screen

func CountLights(width, height int, path string) int {
	return len(process(path, width, height))
}

func DisplayLights(width, height int, path string) {
	screen := process(path, width, height)
	index := utils.NewSet[grid.Coordinate]()
	for _, coordinate := range screen {
		index.Add(coordinate)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if index.Contains(grid.Coordinate{X: x, Y: y}) {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func rect(width, height int) Command {
	return func(screen Screen) Screen {
		var new Screen
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				new = append(new, grid.Coordinate{X: x, Y: y})
			}
		}
		screen = append(screen, new...)
		return utils.Uniq(screen)
	}
}

func row(row, width, amount int) Command {
	return func(screen Screen) Screen {
		var out Screen
		for _, coordinate := range screen {
			if coordinate.Y == row {
				x := (coordinate.X + amount) % width
				out = append(out, grid.Coordinate{X: x, Y: coordinate.Y})
			} else {
				out = append(out, coordinate)
			}

		}
		return out
	}
}

func column(column, height, amount int) Command {
	return func(screen Screen) Screen {
		var out Screen
		for _, coordinate := range screen {
			if coordinate.X == column {
				y := (coordinate.Y + amount) % height
				out = append(out, grid.Coordinate{X: coordinate.X, Y: y})
			} else {
				out = append(out, coordinate)
			}

		}
		return out
	}
}

func process(path string, width, height int) Screen {
	return utils.Reduce(
		parse(path, width, height),
		make(Screen, 0),
		func(screen Screen, next Command) Screen {
			return next(screen)
		})
}

func parse(path string, width, height int) []Command {
	extract := func(line string) Command {
		if strings.Index(line, "rect ") != -1 {
			parts := strings.Split(line[5:], "x")
			assert.True(len(parts) == 2, "expected to parse 2 parts", "parts", parts, "line", line)
			return rect(utils.MustAtoi(parts[0]), utils.MustAtoi(parts[1]))
		} else if strings.Index(line, "rotate column x=") != -1 {
			parts := strings.Split(line[16:], " ")
			assert.True(len(parts) == 3, "expected to parse 3 parts", "parts", parts, "line", line)
			return column(utils.MustAtoi(parts[0]), height, utils.MustAtoi(parts[2]))
		} else if strings.Index(line, "rotate row y=") != -1 {
			parts := strings.Split(line[13:], " ")
			assert.True(len(parts) == 3, "expected to parse 3 parts", "parts", parts, "line", line)
			return row(utils.MustAtoi(parts[0]), width, utils.MustAtoi(parts[2]))
		} else {
			assert.Fail("unable to parse line", "line", line)
			return rect(1, 1)
		}
	}
	return utils.Map(utils.MustReadInput(path), extract)
}
