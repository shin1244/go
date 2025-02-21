package sort

import (
	"runtime"
	"sync"
)

type QuickSort struct {
	MaxThreads int
}

func (q *QuickSort) Sort(arr []int) []int {
	if len(arr) <= 50 {
		return insertionSort(arr)
	}
	idx := partition(arr)
	return append(q.Sort(arr[:idx]), q.Sort(arr[idx+1:])...)
}

// func (q *QuickSort) SortUseThread(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	pivotIdx := partition(arr)
// 	var left, right []int
// 	var wg sync.WaitGroup

// 	if runtime.NumGoroutine() < q.MaxThreads {
// 		wg.Add(2)
// 		go func() {
// 			left = q.SortUseThread(arr[:pivotIdx])
// 			wg.Done()
// 		}()
// 		go func() {
// 			right = q.SortUseThread(arr[pivotIdx+1:])
// 			wg.Done()
// 		}()
// 		wg.Wait()
// 		return append(left, append(arr[pivotIdx:pivotIdx+1], right...)...)
// 	}
// 	return append(q.Sort(arr[:pivotIdx]), q.Sort(arr[pivotIdx+1:])...)
// }

func (q *QuickSort) SortUseThread(arr []int) []int {
	if len(arr) <= 50 {
		return insertionSort(arr)
	}

	pivotIdx := partition(arr)
	var left, right []int
	var wg sync.WaitGroup

	// 쓰레드 개수 제한
	if runtime.NumGoroutine() < runtime.NumCPU() {
		wg.Add(2)
		go func() {
			defer wg.Done()
			left = q.SortUseThread(arr[:pivotIdx])
		}()
		go func() {
			defer wg.Done()
			right = q.SortUseThread(arr[pivotIdx+1:])
		}()
		wg.Wait()
	} else {
		left = q.SortUseThread(arr[:pivotIdx])
		right = q.SortUseThread(arr[pivotIdx+1:])
	}

	return append(left, append(arr[pivotIdx:pivotIdx+1], right...)...)
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func partition(arr []int) int {
	pivot := arr[0]
	i := 1
	j := len(arr) - 1
	for {
		for i < len(arr) && arr[i] < pivot {
			i++
		}
		for j > 0 && arr[j] > pivot {
			j--
		}
		if i >= j {
			arr[0], arr[i-1] = arr[i-1], arr[0]
			return i - 1
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}
