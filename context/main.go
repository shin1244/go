package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())

	go printEverySecond(ctx)
	time.Sleep(time.Second * 5)
	cancel()

	wg.Wait()
}

func printEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-ctx.Done():
			wg.Done()
			return
		}
	}
}
