package main

import "fmt"

func main() {
	go sub()
	for {
		fmt.Println("main loop")
	}
}

func sub() {
	for {
		fmt.Println("Sub loop")
	}
}
