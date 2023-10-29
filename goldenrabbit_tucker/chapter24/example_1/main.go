package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("내 컴퓨터의 고루틴 생성 가능한 수는: %d\n", runtime.GOMAXPROCS(0))
}
