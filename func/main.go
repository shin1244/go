package main

import "fmt"

func sum[T ~int | ~float64](i ...T) T {
	var sum T
	for _, v := range i {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(100, 200))
	fmt.Println(sum(0.3, 0.7))
}
