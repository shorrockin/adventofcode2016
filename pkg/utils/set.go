package utils

type Set[T comparable] map[T]bool

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func NewSetFrom[T comparable](values ...T) Set[T] {
	s := make(Set[T])
	for _, value := range values {
		s[value] = true
	}
	return s
}

func (s Set[T]) Copy() Set[T] {
	return CopyMap(s)
}

func (s Set[T]) Add(value T) {
	s[value] = true
}

func (s Set[T]) AddAll(values ...T) {
	for _, value := range values {
		s[value] = true
	}
}

func (s Set[T]) Remove(value T) {
	delete(s, value)
}

func (s Set[T]) Contains(value T) bool {
	return s[value]
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Values() []T {
	return Keys(s)
}
