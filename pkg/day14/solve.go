package day14

import (
	"crypto/md5"
	"fmt"
)

func Solve(salt string, stretch bool) int {
	hasher := NewHasher(salt, stretch)
	count := 0
	found := 0

	for {
		key := hasher.At(count)

		if index, ok := HasTriple(key); ok {
			for i := 1; i <= 1000; i++ {
				if HasQuintuple(hasher.At(count+i), key[index]) {
					found++
					if found == 64 {
						return count
					}
					break
				}
			}
		}

		count += 1
	}
}

func HasTriple(hash string) (int, bool) {
	for i := 0; i < len(hash)-2; i++ {
		if hash[i] == hash[i+1] && hash[i] == hash[i+2] {
			return i, true
		}
	}
	return 0, false
}

func HasQuintuple(hash string, char byte) bool {
	for i := 0; i < len(hash)-4; i++ {
		if char == hash[i] &&
			hash[i] == hash[i+1] &&
			hash[i] == hash[i+2] &&
			hash[i] == hash[i+3] &&
			hash[i] == hash[i+4] {
			return true
		}
	}
	return false
}

type Hasher struct {
	salt    string
	stretch bool
	cache   map[int]string
}

func NewHasher(salt string, stretch bool) *Hasher {
	return &Hasher{
		salt:    salt,
		stretch: stretch,
		cache:   make(map[int]string),
	}
}

func (h *Hasher) At(index int) string {
	if hash, ok := h.cache[index]; ok {
		return hash
	}

	str := fmt.Sprintf("%s%d", h.salt, index)
	bytes := md5.Sum([]byte(str))
	hash := fmt.Sprintf("%x", bytes)

	if h.stretch {
		for i := 0; i < 2016; i++ {
			bytes = md5.Sum([]byte(hash))
			hash = fmt.Sprintf("%x", bytes)
		}
	}

	h.cache[index] = hash
	return hash
}
