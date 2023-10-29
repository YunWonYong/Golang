package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

var wg sync.WaitGroup

func main() {
	noBufferChannel := make(chan string)
	bufferChannel := make(chan string, 10)

	fmt.Println(len(noBufferChannel), len(bufferChannel))
	wg.Add(2)
	go msgPrint(1, noBufferChannel)
	go msgPrint(2, bufferChannel)

	go pushMsg(noBufferChannel)
	go pushMsg(bufferChannel)

	wg.Wait()
	close(noBufferChannel)
	close(bufferChannel)
}

func pushMsg(buff chan string) {
	for i := 0; i < 10; i++ {
		buff <- uuid.New().String()
	}
	wg.Done()
}

func msgPrint(id int64, buff chan string) {
	for {
		str := <-buff
		fmt.Println(id, str)
	}
}
