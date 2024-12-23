package sort

import "golang.org/x/exp/constraints"

func BinaryInsertSort[T constraints.Ordered](sorted []T, val T) []T {
	idx := findInsert(sorted, val)
	return append(sorted[:idx], append([]T{val}, sorted[idx:]...)...)
}

func findInsert[T constraints.Ordered](arr []T, val T) int {
	if len(arr) == 0 {
		return 0
	}
	mid := len(arr) / 2
	if arr[mid] < val {
		return findInsert(arr[mid+1:], val) + mid + 1
	} else {
		return findInsert(arr[:mid], val)
	}
}
