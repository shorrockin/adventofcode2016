package slices

import (
	"adventofcode2016/pkg/utils/assert"
	"adventofcode2016/pkg/utils/collections"

	"golang.org/x/exp/constraints"
)

func Copy[T any](original []T) []T {
	copied := make([]T, len(original))
	copy(copied, original)
	return copied
}

func Equals[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func Remove[T comparable](slice []T, item T) []T {
	for i, v := range slice {
		if v == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func Intersection[T comparable](left, right []T) []T {
	result := []T{}
	set := collections.NewSetFrom(left...)

	for _, item := range right {
		if set.Contains(item) {
			result = append(result, item)
		}
	}

	return result
}

func Map[T any, K any](data []T, mapper func(T) K) []K {
	result := make([]K, len(data))
	for index, element := range data {
		result[index] = mapper(element)
	}
	return result
}

func MaybeMap[T any, K any](data []T, mapper func(T) (K, bool)) []K {
	result := make([]K, 0, len(data))
	for _, element := range data {
		value, ok := mapper(element)
		if ok {
			result = append(result, value)
		}
	}
	return result
}

func Uniq[T comparable](input []T) []T {
	seen := make(map[T]bool)
	var result []T

	for _, value := range input {
		if _, exists := seen[value]; !exists {
			result = append(result, value)
			seen[value] = true
		}
	}

	return result
}

func Max[T constraints.Ordered](input []T) T {
	if len(input) == 0 {
		var zero T
		return zero
	}

	best := input[0]
	for _, current := range input {
		if current > best {
			best = current
		}
	}

	return best
}

func Count[T any](input []T, selector func(T) bool) int {
	count := 0
	for _, value := range input {
		if selector(value) {
			count++
		}
	}
	return count
}

func Any[T any](input []T, selector func(T) bool) bool {
	for _, value := range input {
		if selector(value) {
			return true
		}
	}
	return false
}

func Filter[T any](input []T, selector func(T) bool) []T {
	out := make([]T, 0)
	for _, value := range input {
		if selector(value) {
			out = append(out, value)
		}
	}

	return out
}

func Reduce[T any, K any](input []T, initial K, reducer func(current K, next T) K) K {
	for _, value := range input {
		initial = reducer(initial, value)
	}
	return initial
}

func Chunk[T any](input []T, size int) [][]T {
	assert.Assert(size > 0, "chunk size must be greater than 0")
	chunks := len(input) / size
	if len(input)%chunks != 0 {
		chunks++
	}

	out := make([][]T, 0, chunks)
	for i := 0; i < chunks; i++ {
		last := (i + 1) * size
		if last > len(input) {
			last = len(input)

		}
		out = append(out, input[i*size:last:last])
	}
	return out
}

func Find[T any](input []T, selector func(T) bool) T {
	for _, value := range input {
		if selector(value) {
			return value
		}
	}
	var zero T
	return zero
}
