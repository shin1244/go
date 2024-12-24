package mymap

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Element[Tkey constraints.Ordered, Tval any] struct {
	key Tkey
	val Tval
}

type SortedMap[Tkey constraints.Ordered, Tval any] struct {
	arr []Element[Tkey, Tval]
}

func (s *SortedMap[Tkey, Tval]) Add(key Tkey, val Tval) {
	idx := insertSort(key, s.arr)
	if len(s.arr) > idx && s.arr[idx].key == key {
		s.arr[idx].val = val
		return
	}
	s.arr = append(s.arr[:idx], append([]Element[Tkey, Tval]{{key, val}}, s.arr...)...)
}

func insertSort[Tkey constraints.Ordered, Tval any](key Tkey, elem []Element[Tkey, Tval]) int {
	if len(elem) == 0 {
		return 0
	}
	mid := len(elem) / 2
	if elem[mid].key < key {
		return insertSort(key, elem[mid+1:]) + mid + 1
	} else {
		return insertSort(key, elem[:mid])
	}
}

func (s *SortedMap[Tkey, Tval]) Get(key Tkey) (Tval, int) {
	for idx, elem := range s.arr {
		if elem.key == key {
			return elem.val, idx
		}
	}

	var temp Element[Tkey, Tval]
	return temp.val, -1
}

func (s *SortedMap[Tkey, Tval]) Remove(key Tkey) bool {
	idx := sort.Search(len(s.arr), func(i int) bool {
		return s.arr[i].key >= key
	})

	if len(s.arr) > idx && s.arr[idx].key == key {
		s.arr = append(s.arr[:idx], s.arr[idx+1:]...)
		return true
	}
	return false
}
