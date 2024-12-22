package sort

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(100)
	}

	assert.False(t, IsSorted(arr), arr)
	QuickSort(arr)
	assert.True(t, IsSorted(arr))
}
