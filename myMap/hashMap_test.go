package mymap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMap(t *testing.T) {
	var h HashMap[int]
	h.Add("A", 100)

	val, ok := h.Get("A")
	assert.True(t, ok)
	assert.Equal(t, 100, val)

	h.Add("B", 200)

	val, ok = h.Get("B")
	assert.True(t, ok)
	assert.Equal(t, 200, val)

	val, ok = h.Get("A")
	assert.True(t, ok)
	assert.Equal(t, 100, val)

	h.Add("C", 300)

	val, ok = h.Get("C")
	assert.True(t, ok)
	assert.Equal(t, 300, val)
}

func TestGoBasicMap(t *testing.T) {
	m := make(map[string]int)

	m["A"] = 100
	m["B"] = 200
	m["C"] = 300

	assert.Equal(t, m["A"], 100)
	assert.Equal(t, m["B"], 200)
	assert.Equal(t, m["C"], 300)
	assert.Equal(t, m["D"], 0)

	_, ok := m["D"]
	assert.False(t, ok)

	delete(m, "A")
	_, ok = m["A"]
	assert.False(t, ok)
}
