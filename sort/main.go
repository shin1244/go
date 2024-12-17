package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 7, 6, 4, 3, 8, 9, 0, 1, 2, 3, 1, 2, 5, 4, 2, 9, 7, 2, 5, 4}

	var count [10]int

	for _, val := range arr {
		count[val]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	result := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		result[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	fmt.Println(count)
	fmt.Println(result)
}
