package main

import (
	"fmt"
)

func main() {
	// slice := []int{}

	// var capacity int
	// for i := 0; i < 10000; i++ {
	// 	slice = append(slice, i)
	// 	if capacity != cap(slice) {
	// 		capacity = cap(slice)
	// 		fmt.Printf("현재 시작 좌표 값: %p, 슬라이스의 길이: %d, 총 용량: %d\n", slice, len(slice), cap(slice))
	// 	}
	// }

	slice := []int{1, 2, 3, 4, 5}
	array := [5]int{1, 2, 3, 4, 5}

	fmt.Println(&slice[0])
	fmt.Println(&array[0])

	changeSlice(slice)
	changeArray(array)

	fmt.Println(slice)
	fmt.Println(array)
}

func changeSlice(slice []int) {
	fmt.Println(&slice[0])
	slice[0] = 100
}

func changeArray(array [5]int) {
	fmt.Println(&array[0])
	array[0] = 100
}
