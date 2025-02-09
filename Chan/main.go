package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int, 2)

	go square3(&wg, ch)

	for i := 1; i < 10; i++ {
		ch <- i
	}

	wg.Wait()
}

func square1() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep")
	}
}

func square2(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Println(n * n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func square3(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminate")
			wg.Done()
			return
		case n := <-ch:
			fmt.Println(n)
			time.Sleep(time.Second)
		}
	}
}
