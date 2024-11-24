package day07

import (
	"adventofcode2016/pkg/assert"
	"adventofcode2016/pkg/utils"
	"slices"
)

const OPENING = '['
const CLOSING = ']'

func PartOne(path string) int {
	lines := utils.MustReadInput(path)
	return len(utils.Filter(lines, SupportsTLS))
}

func PartTwo(path string) int {
	lines := utils.MustReadInput(path)
	return len(utils.Filter(lines, SupportsSSL))
}

func SupportsTLS(ip string) bool {
	return Supportability(ip, false)
}

func SupportsSSL(ip string) bool {
	return Supportability(ip, true)
}

func Supportability(ip string, testSSL bool) bool {
	inHypernet := false
	hasAbba := false
	hasAbbaInHypernet := false
	abas := make([]string, 0)
	babs := make([]string, 0)

	for idx, char := range ip {
		switch char {
		case OPENING:
			assert.False(inHypernet, "hit opening bracket while already open", "ip", ip)
			inHypernet = true
		case CLOSING:
			assert.True(inHypernet, "hit closing bracket while not ignoring", "ip", ip)
			inHypernet = false
		default:
			if idx >= 3 && ip[idx] == ip[idx-3] && ip[idx-1] == ip[idx-2] && ip[idx] != ip[idx-1] {
				switch inHypernet {
				case true:
					hasAbbaInHypernet = true
				case false:
					hasAbba = true
				}
			}

			if idx >= 2 && ip[idx] == ip[idx-2] && ip[idx] != ip[idx-1] {
				switch inHypernet {
				case true:
					babs = append(babs, ip[idx-2:idx+1])
				case false:
					abas = append(abas, ip[idx-2:idx+1])
				}
			}
		}
	}

	if testSSL {
		for _, aba := range abas {
			match := []byte{aba[1], aba[0], aba[1]}
			if slices.Index(babs, string(match)) != -1 {
				return true
			}

		}
		return false
	}

	return hasAbba && !hasAbbaInHypernet
}
