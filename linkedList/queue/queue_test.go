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
