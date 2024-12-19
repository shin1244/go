package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	q := New[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	assert.Equal(t, q.Pop(), 1)
	assert.Equal(t, q.Pop(), 2)
	assert.Equal(t, q.Pop(), 3)
}

func TestPush2(t *testing.T) {
	q := NewSliceQueue[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	assert.Equal(t, q.Pop(), 1)
	assert.Equal(t, q.Pop(), 2)
	assert.Equal(t, q.Pop(), 3)
}

func BenchmarkLinkedListQueue(b *testing.B) {
	s := New[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func BenchmarkSliceQueue(b *testing.B) {
	s := NewSliceQueue[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
