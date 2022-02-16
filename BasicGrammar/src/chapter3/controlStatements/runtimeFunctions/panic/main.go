package main

import "fmt"

func main() {
	defer fmt.Println("Java")
	panic("내가 좋아하는 언어는?")
	defer fmt.Println("JavaScript")
	fmt.Println("Go")
}
