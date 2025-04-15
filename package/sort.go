package test

import (
	"golang.org/x/exp/constraints"
)

type OrderedSlice[T constraints.Ordered] []T

func (o OrderedSlice[T]) Len() int           { return len(o) }
func (o OrderedSlice[T]) Less(i, j int) bool { return o[i] < o[j] }
func (o OrderedSlice[T]) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
