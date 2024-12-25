package maps

import (
	"golang.org/x/exp/constraints"
)

func MaxValue[K comparable, V constraints.Ordered](input map[K]V) V {
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

func Filter[K comparable, V any](input map[K]V, selector func(key K, value V) bool) map[K]V {
	out := make(map[K]V)
	for key, value := range input {
		if selector(key, value) {
			out[key] = value
		}
	}
	return out
}

func Copy[K comparable, V any](original map[K]V) map[K]V {
	copy := make(map[K]V, len(original))
	for key, value := range original {
		copy[key] = value
	}
	return copy
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}
