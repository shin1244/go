package sort

import "golang.org/x/exp/constraints"

func QuickSort[T constraints.Ordered](arr []T) {
	if len(arr) <= 1 {
		return
	}
	idx := patition(arr)
	QuickSort(arr[:idx])
	QuickSort(arr[idx+1:])
}

func IsSorted[T constraints.Ordered](arr []T) bool {
	if len(arr) <= 1 {
		return true
	}
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}
