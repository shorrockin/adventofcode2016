package collections

import "container/list"

// double ended queue
type Dequeue[T any] struct {
	list     *list.List
	popFront bool
}

func NewQueue[T any]() *Dequeue[T] {
	return &Dequeue[T]{list: list.New(), popFront: true}
}

func NewStack[T any]() *Dequeue[T] {
	return &Dequeue[T]{list: list.New(), popFront: false}
}

func (q *Dequeue[T]) Push(value T) {
	q.list.PushBack(value)
}

func (q *Dequeue[T]) PushAll(values ...T) {
	for _, value := range values {
		q.Push(value)
	}
}

func (q *Dequeue[T]) Pop() (T, bool) {
	if q.list.Len() == 0 {
		var zero T
		return zero, false
	}
	var element *list.Element

	if q.popFront {
		element = q.list.Front()
	} else {
		element = q.list.Back()
	}

	q.list.Remove(element)
	return element.Value.(T), true
}

func (q *Dequeue[T]) MustPop() T {
	value, ok := q.Pop()
	if !ok {
		panic("failed to dequeue, expected value to be on queue")
	}
	return value
}

func (q *Dequeue[T]) IsEmpty() bool {
	return q.list.Len() == 0
}

func (q *Dequeue[T]) Len() int {
	return q.list.Len()
}

func (q *Dequeue[T]) Clear() {
	q.list.Init()
}
