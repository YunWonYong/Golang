package main

import (
	"fmt"
)

func main() {
	slice1 := []string{ "A", "B", "C" }
	slice2 := slice1[1: 2]
	fmt.Printf("slice1: %#v, %p\n", slice1, slice1)
	fmt.Printf("slice2: %#v, %p\n", slice2, slice2)

	fmt.Println("slice2[0] = \"Z\"")
	slice2[0] = "Z"

	fmt.Printf("slice1: %#v, %p\n", slice1, slice1)
	fmt.Printf("slice2: %#v, %p\n", slice2, slice2)
}
