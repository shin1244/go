package main

import (
	"fmt"
	"time"
)

func main() {
	wg.Add(10)
	account := &Account{1}
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
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
