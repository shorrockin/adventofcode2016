package day04

import (
	"adventofcode2016/pkg/assert"
	"adventofcode2016/pkg/utils"
	"regexp"
	"slices"
	"sort"
	"strings"
)

var ROOM_REGEXP *regexp.Regexp
var ALPHABET []rune

func init() {
	ROOM_REGEXP = regexp.MustCompile(`([a-z\-]+)\-(\d+)\[([a-z]{5})\]`)
	ALPHABET = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
}

type Room struct {
	name     string
	sector   int
	checksum string
}

func NewRoom(value string) Room {
	matches := ROOM_REGEXP.FindStringSubmatch(value)
	assert.NotNil(matches, "expected value string to match regexp", "value", value)
	assert.True(len(matches) == 4, "expected FindAllString to return 4", "len(matches)", len(matches), "input", value, "regexp", ROOM_REGEXP)
	return Room{matches[1], utils.MustAtoi(matches[2]), matches[3]}
}

func (r Room) Decrypt() string {
	var builder strings.Builder
	for _, char := range r.name {
		if char == '-' {
			builder.WriteRune(' ')
		} else {
			index := slices.Index(ALPHABET, char)
			builder.WriteRune(ALPHABET[(index+r.sector)%len(ALPHABET)])
		}
	}
	return builder.String()
}

func (r Room) GenerateChecksum() string {
	frequencies := make(map[rune]int)
	for _, char := range r.name {
		if char == '-' {
			continue
		}
		frequencies[char] += 1
	}
	keys := make([]rune, 0, len(frequencies))
	for k := range frequencies {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		left := keys[i]
		right := keys[j]
		countLeft := frequencies[left]
		countRight := frequencies[right]
		if countLeft == countRight {
			return left < right
		}
		return countLeft > countRight
	})
	return string(keys[:5])
}

func SumSectors(path string) int {
	rooms := utils.Map(utils.MustReadInput(path), NewRoom)
	rooms = utils.Filter(rooms, func(room Room) bool {
		return room.checksum == room.GenerateChecksum()
	})
	return utils.Reduce(rooms, 0, func(accum int, room Room) int {
		return accum + room.sector
	})
}

func FindNorthPoleSector(path string) int {
	rooms := utils.Map(utils.MustReadInput(path), NewRoom)
	for _, room := range rooms {
		if room.Decrypt() == "northpole object storage" {
			return room.sector
		}
	}
	return 0
}
