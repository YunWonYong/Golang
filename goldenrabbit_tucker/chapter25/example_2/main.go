package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	noBufferChannel := make(chan string)
	bufferChannel := make(chan string, 10)
	quitChannel := make(chan int)
	quitNum := 0
	go pushMsg(noBufferChannel, quitChannel)
	go pushMsg(bufferChannel, quitChannel)

	for {
		select {
		case str := <-noBufferChannel:
			fmt.Println(1, str)
		case str := <-bufferChannel:
			fmt.Println(2, str)
		case n := <-quitChannel:
			quitNum += n
			if quitNum == 2 {
				close(noBufferChannel)
				close(bufferChannel)
				return
			}
		}
	}
}

func pushMsg(buff chan string, quit chan int) {
	for i := 0; i < 10; i++ {
		buff <- uuid.New().String()
	}
	quit <- 1
}
