package main

import (
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Account struct {
	Balance int
}

func main() {
	account := new(Account)
	routineCount := runtime.GOMAXPROCS(0)
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		go func() {
			for {
				if account.Balance < 0 {
					panic("잔고가 0보다 작습니다.")
				}
				account.Balance += 1000
				time.Sleep(time.Millisecond)
				account.Balance -= 1000
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
