package singlelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.Root)
	assert.Nil(t, l.Tail)

	l.PushBack(1)

	assert.NotNil(t, l.Root)
	assert.Equal(t, 1, l.Root.Value)
	assert.Equal(t, 1, l.Tail.Value)

	l.PushBack(2)

	assert.NotNil(t, l.Root)
	assert.Equal(t, 1, l.Root.Value)
	assert.Equal(t, 2, l.Tail.Value)

	l.PushBack(3)

	assert.NotNil(t, l.Root)
	assert.Equal(t, 1, l.Root.Value)
	assert.Equal(t, 3, l.Tail.Value)

	assert.Equal(t, 3, l.Count())
}

func TestPushFront(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.Root)
	assert.Nil(t, l.Tail)

	l.PushFront(1)

	assert.NotNil(t, l.Root)
	assert.Equal(t, 1, l.Root.Value)
	assert.Equal(t, 1, l.Tail.Value)

	l.PushFront(2)

	assert.NotNil(t, l.Root)
	assert.Equal(t, 2, l.Root.Value)
	assert.Equal(t, 1, l.Tail.Value)

	l.PushFront(3)

	assert.NotNil(t, l.Root)
	assert.Equal(t, 3, l.Root.Value)
	assert.Equal(t, 1, l.Tail.Value)

	assert.Equal(t, 3, l.Count())

	getNode := l.GetAt(1)
	l.InsertAfter(getNode, 0)

}

func TestLinkedList(t *testing.T) {
	var l LinkedList[int]

	l.PushFront(1)
	l.PushFront(0)
	l.PushBack(2)

	assert.Equal(t, 3, l.Count())

	targetNode := l.GetAt(1)
	l.InsertAfter(targetNode, 5)

	assert.Equal(t, []any{0, 1, 5, 2}, l.ListToArray())

	targetNode = l.GetAt(2)
	l.InsertBefore(targetNode, 4)
	assert.Equal(t, []any{0, 1, 4, 5, 2}, l.ListToArray())

	l.Reverse()
	t.Log(l.ListToArray())
}
