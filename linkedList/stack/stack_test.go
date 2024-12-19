package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, s.Pop(), 3)
	assert.Equal(t, s.Pop(), 2)
	assert.Equal(t, s.Pop(), 1)
}

func TestPush2(t *testing.T) {
	s := NewSliceStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, s.Pop(), 3)
	assert.Equal(t, s.Pop(), 2)
	assert.Equal(t, s.Pop(), 1)
}

func BenchmarkLinkedListStack(b *testing.B) {
	s := New[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func BenchmarkSliceStack(b *testing.B) {
	s := NewSliceStack[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
