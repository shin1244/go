package sort

import (
	"fmt"
	"math/rand/v2"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	arr := rand.Perm(10000)
	arr = Sort(arr, &MergeSort{})
	for i := 0; i < len(arr)-1; i++ {
		assert.True(t, arr[i] <= arr[i+1])
	}
	fmt.Println(arr)
}
func TestQuickSort(t *testing.T) {
	arr := rand.Perm(10000)
	arr = Sort(arr, &QuickSort{})
	for i := 0; i < len(arr)-1; i++ {
		assert.True(t, arr[i] <= arr[i+1])
	}
	fmt.Println(arr)
}

func TestSortUseThread(t *testing.T) {
	quickSort := &QuickSort{
		MaxThreads: runtime.NumCPU(),
	}
	fmt.Println(runtime.NumCPU())
	arr := rand.Perm(10000)
	arr = SortUseThread(arr, quickSort)
	for i := 0; i < len(arr)-1; i++ {
		assert.True(t, arr[i] <= arr[i+1])
	}
}

func BenchmarkQuickSort(b *testing.B) {
	originalArr := rand.Perm(10000)
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(originalArr))
		copy(arr, originalArr)
		Sort(arr, &QuickSort{})
	}
}

func BenchmarkQuickSortUseThread(b *testing.B) {
	quickSort := &QuickSort{
		MaxThreads: runtime.NumCPU(),
	}
	originalArr := rand.Perm(10000)
	for i := 0; i < b.N; i++ {
		arr := make([]int, len(originalArr))
		copy(arr, originalArr)
		SortUseThread(arr, quickSort)
	}
}
