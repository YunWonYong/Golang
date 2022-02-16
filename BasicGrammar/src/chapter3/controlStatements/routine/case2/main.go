package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Printf("Version: %s\n", runtime.Version())

	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("NumCgoCall: %d\n", runtime.NumCgoCall())
	go fmt.Println("go routine + 1")
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("NumCgoCall: %d\n", runtime.NumCgoCall())
	go func() {
		fmt.Println("go routine 2")
		for x := 0; x < 100; x++ {
			fmt.Printf("go routine x:%d\n", x)
		}
	}()
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("NumCgoCall: %d\n", runtime.NumCgoCall())

	go func() {
		fmt.Println("go routine 3")
		for x := 0; x < 100; x++ {
			fmt.Printf("go routine x:%d\n", x)
		}
	}()
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("NumCgoCall: %d\n", runtime.NumCgoCall())

	defer fmt.Print("go routine은 마지막에 실행해서??") //이거 없으면 go routine으로 출력하는 내용 없음
}
