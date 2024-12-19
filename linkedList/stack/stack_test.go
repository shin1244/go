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
