package main

import "fmt"

func main() {
	x := 0
	go func() {
		for x < 100 {
			x++
			fmt.Printf("go routine x:= %d\n", x)
		}
	}()
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
			fmt.Println("main dead")
		}
	}()
	for {
		if x == 100 {
			panic("go routine dead")
		}
	}
}
