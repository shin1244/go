package test

import "golang.org/x/exp/constraints"

type Heap[T constraints.Ordered] []T

func (h Heap[T]) Len() int           { return len(h) }
func (h Heap[T]) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap[T]) Push(x any)        { *h = append(*h, x.(T)) }
func (h *Heap[T]) Pop() any {
	old := *h
	element := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return element
}
