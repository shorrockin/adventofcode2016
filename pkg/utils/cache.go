package utils

type Cache[K comparable, V any] map[K]V

func (c Cache[K, V]) Memoize(key K, getter func() V) V {
	value, ok := c[key]
	if !ok {
		value = getter()
		c[key] = value
	}
	return value
}

func NewCache[K comparable, V any]() Cache[K, V] {
	return make(Cache[K, V])
}
