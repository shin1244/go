package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Sum(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("Sum of %d to %d is %d\n", a, b, sum)
	wg.Done()
}
