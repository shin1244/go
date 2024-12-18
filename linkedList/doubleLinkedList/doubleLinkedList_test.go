package doublelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushBack(1)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.root.Value)
	assert.Equal(t, 1, l.tail.Value)

	l.PushBack(2)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.root.Value)
	assert.Equal(t, 2, l.tail.Value)

	l.PushBack(3)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.root.Value)
	assert.Equal(t, 3, l.tail.Value)

	assert.Equal(t, 3, l.Count())
}

func TestPushFront(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushFront(1)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.root.Value)
	assert.Equal(t, 1, l.tail.Value)

	l.PushFront(2)

	assert.NotNil(t, l.root)
	assert.Equal(t, 2, l.root.Value)
	assert.Equal(t, 1, l.tail.Value)

	l.PushFront(3)

	assert.NotNil(t, l.root)
	assert.Equal(t, 3, l.root.Value)
	assert.Equal(t, 1, l.tail.Value)

	assert.Equal(t, 3, l.Count())

	getNode := l.GetAt(2)
	l.InsertAfter(getNode, 0)

	getNode = l.GetAt(0)
	l.InsertBefore(getNode, 0)

	getNode = l.GetAt(2)
	l.InsertBefore(getNode, 5)
	l.InsertAfter(getNode, 5)

	t.Log(l.ListToArray(), l.count)

	t.Log(l.PopFront())
	t.Log(l.PopBack())
	t.Log(l.PopBack())
	t.Log(l.PopFront())
	t.Log(l.PopFront())
	t.Log(l.PopFront())
	t.Log(l.PopFront())
	t.Log(l.ListToArray(), l.count)
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
	assert.Equal(t, []any{2, 5, 4, 1, 0}, l.ListToArray())

	// targetNode = l.GetAt(4)
	// l.Remove(targetNode)
	// assert.Equal(t, []any{0, 1, 4, 5}, l.ListToArray())

	// assert.Equal(t, 0, l.PopFront().Value)
	// assert.Equal(t, []any{1, 4, 5}, l.ListToArray())

	// targetNode = l.GetAt(1)
	// l.Remove(targetNode)
	// assert.Equal(t, []any{1, 5}, l.ListToArray())

	// targetNode = l.GetAt(1)
	// l.Remove(targetNode)
	// assert.Equal(t, []any{1}, l.ListToArray())

	// targetNode = l.GetAt(0)
	// l.Remove(targetNode)
	// assert.Equal(t, []any{}, l.ListToArray())
}
