package main

import "fmt"

func main() {
	code1()
}

func code1() {
	var a int
	var p *int

	fmt.Printf("p variable default value: %v\n", p)
	// fmt.Printf("p variable default memoey value: %v\n", *p) // panic

	a = 10
	p = &a

	fmt.Printf("p variable value: %p\n", p)
	fmt.Printf("p variable memory value: %d\n", *p)

	fmt.Printf("a variable value: %d\n", a)
	fmt.Printf("a variable memory address: %p\n", &a)

	*p = 100

	fmt.Println("*p = 100")

	fmt.Printf("p variable value: %p\n", p)
	fmt.Printf("p variable memory value: %d\n", *p)

	fmt.Printf("a variable value: %d\n", a)
	fmt.Printf("a variable memory address: %p\n", &a)
}
