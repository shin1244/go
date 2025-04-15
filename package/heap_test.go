package test

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	// 정수형 힙 테스트
	h := &Heap[int]{}
	heap.Init(h)

	// 힙에 요소 추가
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	for _, v := range values {
		heap.Push(h, v)
	}

	// 힙에서 요소를 꺼내며 정렬된 순서인지 확인
	expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
	for i, want := range expected {
		got := heap.Pop(h).(int)
		if got != want {
			t.Errorf("힙 Pop 실패: 인덱스 %d에서 %d를 기대했으나 %d를 얻음", i, want, got)
		}
	}
}
