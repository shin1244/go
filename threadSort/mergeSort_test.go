package mergeSort

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	arr := rand.Perm(1000)
	arr = MergeSort(arr)
	for i := 0; i < len(arr)-1; i++ {
		assert.True(t, arr[i] <= arr[i+1])
	}
	fmt.Println(arr)
}

func TestMergeSortUseThread(t *testing.T) {
	arr := rand.Perm(1000)
	arr = MergeSortUseThread(arr)
	for i := 0; i < len(arr)-1; i++ {
		assert.True(t, arr[i] <= arr[i+1])
	}
	fmt.Println(arr)
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := rand.Perm(10000000)
		MergeSort(arr)
	}
}

func BenchmarkMergeSortUseThread(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := rand.Perm(10000000)
		MergeSortUseThread(arr)
	}
}
