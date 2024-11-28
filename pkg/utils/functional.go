package utils

import "golang.org/x/exp/constraints"

func Map[T any, K any](data []T, mapper func(T) K) []K {
	result := make([]K, len(data))
	for index, element := range data {
		result[index] = mapper(element)
	}
	return result
}

func MapConditional[T any, K any](data []T, mapper func(T) (K, bool)) []K {
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
	seen := NewSet[T]()
	var result []T

	for _, value := range input {
		if !seen.Contains(value) {
			result = append(result, value)
			seen.Add(value)
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

func MaxMapValue[K comparable, V constraints.Ordered](input map[K]V) V {
	var zero V
	if len(input) == 0 {
		return zero
	}

	best := zero
	for _, value := range input {
		if value > best {
			best = value
		}
	}

	return best
}

func MaxValue[T constraints.Ordered](left T, right T) T {
	if left > right {
		return left
	}
	return right
}

func MinValue[T constraints.Ordered](left T, right T) T {
	if left < right {
		return left
	}
	return right
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

func FilterMap[K comparable, V any](input map[K]V, selector func(key K, value V) bool) map[K]V {
	out := make(map[K]V)
	for key, value := range input {
		if selector(key, value) {
			out[key] = value
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
