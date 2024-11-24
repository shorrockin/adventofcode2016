package day05

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func PartOne(id []byte) string {
	index := 0
	var builder strings.Builder

	for {
		hashed := Hash(id, index)
		// general speed up over string comparison, 15 == 0f
		if hashed[0] == '\x00' && hashed[1] == '\x00' && hashed[2] <= 15 {
			builder.WriteString(fmt.Sprintf("%x", hashed[2]))
			if builder.Len() == 8 {
				return builder.String()
			}
		}

		index++
	}
}

func PartTwo(id []byte) string {
	index := 0
	result := "        "

	for {
		hashed := Hash(id, index)
		index++

		if hashed[0] == '\x00' && hashed[1] == '\x00' && hashed[2] <= 15 {
			str := fmt.Sprintf("%x", hashed)
			pos, err := strconv.Atoi(string(str[5]))
			if err != nil || pos > 7 || result[pos] != ' ' {
				continue
			}

			updated := []rune(result)
			updated[pos] = rune(str[6])
			result = string(updated)
			if strings.Index(result, " ") == -1 {
				return result
			}
		}
	}
}

func Hash(bytes []byte, index int) [16]byte {
	bytes = append(bytes, []byte(strconv.Itoa(index))...)
	return md5.Sum(bytes)
}
