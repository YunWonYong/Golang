package main

import "fmt"

func main() {
	const A = 1     // Untyped const
	const B int = 1 // typed const
	var c int64 = 1
	fmt.Printf("A: %v, B: %v\n", A, B)
	fmt.Printf("A: %T, B: %T\n", A, B)
	fmt.Println(A + B)
	fmt.Println(c + A)
	//fmt.Println(c + B) error
}
