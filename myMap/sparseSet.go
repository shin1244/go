package mymap

import (
	"golang.org/x/exp/constraints"
)

type SparseSet[Tkey constraints.Ordered, Tval any] struct {
	dense  []Element[Tkey, Tval]
	sparse map[Tkey]int
}

func NewSparseSet[Tkey constraints.Ordered, Tval any]() *SparseSet[Tkey, Tval] {
	return &SparseSet[Tkey, Tval]{
		sparse: make(map[Tkey]int),
	}
}

func (s *SparseSet[Tkey, Tval]) Add(Akey Tkey, Aval Tval) {
	if idx, ok := s.sparse[Akey]; ok {
		s.dense[idx].val = Aval
		return
	}
	s.dense = append(s.dense, Element[Tkey, Tval]{
		key: Akey,
		val: Aval,
	})
	s.sparse[Akey] = len(s.dense) - 1
}

func (s *SparseSet[Tkey, Tval]) Get(Gkey Tkey) (Gval Tval, found bool) {
	if idx, ok := s.sparse[Gkey]; ok {
		Gval = s.dense[idx].val
		found = true
		return
	}
	found = false
	return
}

func (s *SparseSet[Tkey, Tval]) Remove(Rkey Tkey) bool {
	if idx, ok := s.sparse[Rkey]; ok {
		last := len(s.dense) - 1
		if idx < last {
			s.dense[idx] = s.dense[last]
			s.sparse[s.dense[last].key] = idx
		}
		s.dense = s.dense[:last]
		delete(s.sparse, Rkey)
		return true
	}
	return false
}

type Iterator[Tkey constraints.Ordered, Tval any] struct {
	dense []Element[Tkey, Tval]
	idx   int
}

func (s *SparseSet[Tkey, Tval]) Iterator() *Iterator[Tkey, Tval] {
	return &Iterator[Tkey, Tval]{
		dense: s.dense,
		idx:   0,
	}
}

func (i *Iterator[Tkey, Tval]) IsEnd() bool {
	return i.idx < len(i.dense)
}

func (i *Iterator[Tkey, Tval]) Next() {
	i.idx++
}

func (i *Iterator[Tkey, Tval]) Get() Element[Tkey, Tval] {
	return i.dense[i.idx]
}
