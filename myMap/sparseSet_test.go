package mymap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSparseSet(t *testing.T) {
	s := NewSparseSet[string, int]()

	s.Add("C", 30)
	s.Add("B", 20)
	s.Add("A", 10)

	v, ok := s.Get("B")
	assert.True(t, ok)
	assert.Equal(t, 20, v)

	v, ok = s.Get("C")
	assert.True(t, ok)
	assert.Equal(t, 30, v)

	v, ok = s.Get("A")
	assert.True(t, ok)
	assert.Equal(t, 10, v)

	_, ok = s.Get("D")
	assert.False(t, ok)

	removed := s.Remove("B")
	assert.True(t, removed)

	_, ok = s.Get("B")
	assert.False(t, ok)

	for i := s.Iterator(); i.IsEnd(); i.Next() {
		elem := i.Get()
		if elem.key == "A" {
			assert.Equal(t, elem.val, 10)
		} else if elem.key == "C" {
			assert.Equal(t, 30, elem.val)
		} else {
			t.Fail()
		}
	}
}
