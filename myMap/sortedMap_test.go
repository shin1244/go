package mymap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedMap(t *testing.T) {
	var s SortedMap[string, int]

	s.Add("A", 10)

	val, idx := s.Get("A")
	assert.Equal(t, val, 10)
	assert.Equal(t, idx, 0)

	s.Add("C", 20)

	val, idx = s.Get("C")
	assert.Equal(t, val, 20)
	assert.Equal(t, idx, 1)

	s.Add("B", 30)

	val, idx = s.Get("B")
	assert.Equal(t, val, 30)
	assert.Equal(t, idx, 1)
}

func TestSortedMapOverlapped(t *testing.T) {
	var s SortedMap[string, int]

	s.Add("A", 10)
	s.Add("A", 100)

	assert.Equal(t, 1, len(s.arr))
}

func TestSortedMapGetEmpty(t *testing.T) {
	var s SortedMap[string, int]

	s.Add("A", 100)

	val, idx := s.Get("B")
	assert.Equal(t, 0, val)
	assert.Equal(t, -1, idx)
}

func TestSortedMapRemove(t *testing.T) {
	var s SortedMap[string, int]

	s.Add("cccc", 30)
	s.Add("bbb", 20)
	s.Add("aaa", 10)

	removed := s.Remove("bbb")
	assert.True(t, removed)
	removed = s.Remove("bbb")
	assert.False(t, removed)

	assert.Equal(t, 2, len(s.arr))

	assert.Equal(t, "aaa", s.arr[0].key)
	assert.Equal(t, "cccc", s.arr[1].key)
}
