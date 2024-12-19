package stack

type SliceStack[T any] struct {
	arr []T
}

func NewSliceStack[T any]() *SliceStack[T] {
	return &SliceStack[T]{}
}

func (s *SliceStack[T]) Push(val T) {
	s.arr = append(s.arr, val)
}

func (s *SliceStack[T]) Pop() T {
	var result T
	if len(s.arr) == 0 {
		return result
	}
	result = s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return result
}
