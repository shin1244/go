package main

import (
	"fmt"
	"time"
)

func main() {
	var jobList [10]job
	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{index: i}
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go jobList[i].Do()

	}
	wg.Wait()
}

func PrintHello() {
	text := []string{"H", "e", "l", "l", "o", " ", "W", "o", "r", "l", "d", "!"}
	for _, v := range text {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%s ", v)
	}
}

func PrintNumber() {
	for i := 0; i < 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
