package day15

import (
	"adventofcode2016/pkg/assert"
	"adventofcode2016/pkg/utils"
	"regexp"
)

func Solve(path string) int {
	discs := parse(path)
	remainders := make([]int, len(discs))
	moduli := make([]int, len(discs))

	for idx, disc := range discs {
		remainders[idx] = disc.positions - ((disc.start + disc.id) % disc.positions)
		moduli[idx] = disc.positions
	}
	return utils.CRT(remainders, moduli)
}

type Disc struct {
	id        int
	positions int
	start     int
}

func parse(path string) []Disc {
	pattern, err := regexp.Compile(`Disc #(\d+) has (\d+) positions; at time=0, it is at position (\d+).`)
	assert.NoError(err, "expected regexp to compile, it did not")

	lines := utils.MustReadInput(path)
	out := make([]Disc, len(lines))

	for idx, line := range lines {
		parts := pattern.FindAllStringSubmatch(line, -1)
		assert.Equal(4, len(parts[0]), "expected 4 parts", "parts", parts[0])
		out[idx] = Disc{
			id:        utils.MustAtoi(string(parts[0][1])),
			positions: utils.MustAtoi(string(parts[0][2])),
			start:     utils.MustAtoi(string(parts[0][3])),
		}

	}
	return out
}
