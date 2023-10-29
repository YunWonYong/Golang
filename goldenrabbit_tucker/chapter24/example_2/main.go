package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for index, at := range "hello world!!!" {
		wg.Add(1)
		go func() {
			fmt.Printf("%d번 서브루틴의 단어: %s\n", index, string(at))
			wg.Done()
		}()
	}
	wg.Wait()
}
