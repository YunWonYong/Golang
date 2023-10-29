package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	wg sync.WaitGroup
)

type Account struct {
	Balance int64
}

func main() {
	routineCount := runtime.GOMAXPROCS(0)
	users := make(map[string]*Account)
	for i := 0; i < routineCount; i++ {
		account := new(Account)
		account.Balance = getRandom(1000)
		users[uuid.New().String()] = account
	}

	wg.Add(routineCount)
	for id, account := range users {
		fmt.Println(id, account)
		go func(id string, account *Account) {
			for {
				if account.Balance <= 0 {
					fmt.Printf("id: %s allin\n", id)
					break
				} else if account.Balance > 5000 {
					fmt.Printf("id: %s, win money: %d\n", id, account.Balance)
					break
				} else if getRandom(10)%2 == 0 {
					account.Balance += getRandom(1000)
					continue
				}
				account.Balance -= getRandom(1000)
			}
			wg.Done()
		}(id, account)
	}
	wg.Wait()
}

func getRandom(base int64) int64 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(base)
}
