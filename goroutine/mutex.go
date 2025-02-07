package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()

	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}

	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func diningProblem(name string, first, second *sync.Mutex) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s is trying to eat.\n", name)
		first.Lock()
		fmt.Printf("%s got %s.\n", name, "first")
		second.Lock()
		fmt.Printf("%s got %s.\n", name, "second")

		fmt.Printf("%s is eating.\n", name)
		time.Sleep(time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}
