package collections

import "container/heap"

type PqNode[T comparable] struct {
	Contents  T
	Priority  float64
	Parent    *PqNode[T]
	PathDepth int
}

func NewPqNode[T comparable](contents T, priority float64, parent *PqNode[T]) *PqNode[T] {
	depth := 0
	if parent != nil {
		depth = parent.PathDepth + 1
	}
	return &PqNode[T]{Contents: contents, Priority: priority, Parent: parent, PathDepth: depth}
}

type Heap[T comparable] []*PqNode[T]

func (pq Heap[T]) Len() int           { return len(pq) }
func (pq Heap[T]) Less(i, j int) bool { return pq[i].Priority < pq[j].Priority }
func (pq Heap[T]) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *Heap[T]) Push(x any) {
	item := x.(*PqNode[T])
	*pq = append(*pq, item)
}

func (pq *Heap[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}

type PriorityQueue[T comparable] struct {
	heap Heap[T]
}

func (pq *PriorityQueue[T]) Push(contents T, priority float64, parent *PqNode[T]) {
	heap.Push(&pq.heap, NewPqNode(contents, priority, parent))
}

func (pq *PriorityQueue[T]) Pop() T {
	return pq.PopNode().Contents
}

func (pq *PriorityQueue[T]) PopNode() *PqNode[T] {
	return heap.Pop(&pq.heap).(*PqNode[T])
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.heap.Len()
}

func NewPriorityQueue[T comparable]() PriorityQueue[T] {
	data := make(Heap[T], 0)
	heap.Init(&data)

	return PriorityQueue[T]{heap: data}
}
