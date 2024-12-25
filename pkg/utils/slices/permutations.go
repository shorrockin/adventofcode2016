package slices

func Permutations[T any](data []T) [][]T {
	var out [][]T
	permute(data, 0, &out, func(data []T, at int) bool { return true })
	return out
}

func PermutationsFiltered[T any](data []T, filter func(data []T, at int) bool) [][]T {
	var out [][]T
	permute(data, 0, &out, filter)
	return out
}

func permute[T any](data []T, start int, result *[][]T, filter func(data []T, at int) bool) {
	// permutation is complete, copy input into result
	if start == len(data)-1 {
		temp := make([]T, len(data))
		copy(temp, data)
		*result = append(*result, temp)
		return
	}

	// iterate over all elements, swapping the specified value
	// then recursing, then backtracking the swap
	for i := start; i < len(data); i++ {
		data[start], data[i] = data[i], data[start] // swap
		if filter(data, start) {
			permute(data, start+1, result, filter)
		}
		data[start], data[i] = data[i], data[start] // backtrack
	}
}
