package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexpExtraction(t *testing.T) {
	assert.Equal(t, Room{"abcd-efg-hij", 123, "abcde"}, NewRoom("abcd-efg-hij-123[abcde]"))
}

func TestChecksumGeneration(t *testing.T) {
	assert.Equal(t, "abcde", Room{"abcd-efg-hij", 123, "abcde"}.GenerateChecksum())
	assert.Equal(t, "abxyz", Room{"aaaaa-bbb-z-y-x", 123, "abxyz"}.GenerateChecksum())
}

func TestPartOne(t *testing.T) {
	assert.Equal(t, 361724, SumSectors("input.txt"))
}

func TestDecryption(t *testing.T) {
	assert.Equal(t, "very encrypted name", Room{"qzmt-zixmtkozy-ivhz", 343, "abcde"}.Decrypt())
}

func TestPartTwo(t *testing.T) {
	assert.Equal(t, 482, FindNorthPoleSector("input.txt"))
}
