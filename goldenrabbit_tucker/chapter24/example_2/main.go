package main

import "fmt"

func main() {
	fmt.Print("hello")
	endFlag := false
	go func() {
		fmt.Println(" world!!!")
		endFlag = true
	}()

	for !endFlag {
	}
}
