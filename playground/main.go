package main

import (
	"fmt"
)

func main() {
	slice := []int{}

	var capacity int
	for i := 0; i < 10000; i++ {
		slice = append(slice, i)
		if capacity != cap(slice) {
			capacity = cap(slice)
			fmt.Printf("현재 시작 좌표 값: %p, 슬라이스의 길이: %d, 총 용량: %d\n", slice, len(slice), cap(slice))
		}
	}
}
