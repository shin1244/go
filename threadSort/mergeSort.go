package sort

import "sync"

type MergeSort struct{}

func (m *MergeSort) Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := m.Sort(arr[:mid])
	right := m.Sort(arr[mid:])

	return merge(left, right)
}

func (m *MergeSort) SortUseThread(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	var left, right []int
	var wg sync.WaitGroup
	wg.Add(2)
	if mid > 10000 {
		go func() {
			left = m.Sort(arr[:mid])
			wg.Done()
		}()
		go func() {
			right = m.Sort(arr[mid:])
			wg.Done()
		}()
		wg.Wait()
	} else {
		left = m.Sort(arr[:mid])
		right = m.Sort(arr[mid:])
	}

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	leftIdx := 0
	rightIdx := 0
	idx := 0

	for len(left) > leftIdx || len(right) > rightIdx {
		var leftMerge bool
		if leftIdx >= len(left) {
			leftMerge = false
		} else if rightIdx >= len(right) {
			leftMerge = true
		} else {
			leftMerge = left[leftIdx] < right[rightIdx]
		}
		if leftMerge {
			result[idx] = left[leftIdx]
			leftIdx++
		} else {
			result[idx] = right[rightIdx]
			rightIdx++
		}
		idx++
	}
	return result
}
