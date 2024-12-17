package day16

func Solve(value string, length int) string {
	runes := make([]rune, 0, length)
	for _, r := range value {
		runes = append(runes, r)
	}

	for len(runes) < length {
		runes = append(runes, '0')
		for i := len(runes) - 2; i >= 0; i-- {
			if runes[i] == '0' {
				runes = append(runes, '1')
			} else {
				runes = append(runes, '0')
			}
		}
	}

	return checksum(runes[:length])
}

func checksum(input []rune) string {
	for {
		result := make([]rune, 0, len(input)/2)
		for i := 0; i < len(input); i += 2 {
			if input[i] == input[i+1] {
				result = append(result, '1')
			} else {
				result = append(result, '0')
			}
		}

		if len(result)%2 != 0 {
			return string(result)
		}

		input = result
	}
}
