package day06

import (
	"adventofcode2016/pkg/utils"
	"strings"
)

func Solve(path string, partOne bool) string {
	lines := utils.MustReadInput(path)
	frequencies := make([]map[rune]int, len(lines[0]))

	for _, line := range lines {
		for index, char := range line {
			if lookup := frequencies[index]; lookup == nil {
				frequencies[index] = make(map[rune]int)
			}
			frequencies[index][char] += 1
		}
	}

	var builder strings.Builder
	for _, lookup := range frequencies {
		var winner rune
		var count int

		for char, current := range lookup {
			if winner != 0 && partOne && current < count {
				continue
			}
			if winner != 0 && !partOne && current > count {
				continue
			}
			winner = char
			count = current

		}
		builder.WriteRune(winner)
	}

	return builder.String()
}
