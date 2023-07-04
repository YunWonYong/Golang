package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	noSeedRandomCheck()
	noLoopSetSeedRandomCheck()
	loopSetSeedRandomCheck()
}

func noSeedRandomCheck() {
	sameCnt := 0
	for i := 0; i < 100; i++ {
		if rand.Intn(100) == rand.Intn(100) {
			sameCnt++
		}
	}

	fmt.Printf("no seed\n random value same count: %d\n", sameCnt)
}

func noLoopSetSeedRandomCheck() {
	rand.Seed(time.Now().UnixNano())

	sameCnt := 0

	for i := 0; i < 100; i++ {
		if rand.Intn(100) == rand.Intn(100) {
			sameCnt++
		}
	}

	fmt.Printf("seed initialization outside the loop\n random value same count: %d\n", sameCnt)
}

func loopSetSeedRandomCheck() {
	sameCnt := 0

	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(100) == rand.Intn(100) {
			sameCnt++
		}
	}

	fmt.Printf("seed initialization inside the loop\n random value same count: %d\n", sameCnt)
}
