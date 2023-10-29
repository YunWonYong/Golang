package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go routine(ctx)

	tick := time.Tick(time.Second * 10)

	for {
		<-tick
		cancel()
		break
	}

	fmt.Println("main routine close")
}

func routine(ctx context.Context) {
	tick := time.Tick(time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-ctx.Done():
			fmt.Println("sub routine close")
			return
		}

	}
}
