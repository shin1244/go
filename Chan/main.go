package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go square()
	ch <- 9
}

func square() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep")
	}
}
