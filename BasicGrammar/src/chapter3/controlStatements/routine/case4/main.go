package main

import (
	"fmt"
)

var (
	x            = 0
	routineCount = 0
)

func main() {
	go goRoutineRun("1번")
	go goRoutineRun("2번")
	go goRoutineRun("3번")
	go func() {
		goRoutineRun("4번")
	}()

	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	for {
		if routineCount == 0 {
			panic("main")
		}
	}
}

func goRoutineRun(name string) {
	routineCount++
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("먼저 끝난 루틴: %s\n", x)
			routineCount--
		}
	}()
	for {
		x++
		if x > 100000 {
			panic(name)
		}
	}
}
