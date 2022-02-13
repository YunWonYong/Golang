package main

import "fmt"

func main() {
	diviserFn := func(x, y int) int { return x % y }
	fmt.Printf("diviserFn type: %T\n", diviserFn)
	var plus func(x, y int) int = func(x, y int) int { return x + y }
	fmt.Printf("plus func memory address: %p\n", plus)
	fmt.Println(plus(5, 10))
	//fmt.Printf("anonymous func : %#v\n", func(x, y int) int { return x - y })
	fmt.Printf("anonymous func call result: %#v\n", func(x, y int) int { return x - y }(10, 5) /*IIFE[Immediately Invoked Function Expressions]*/)
	var multiply = func(x, y int) int { return x * y }
	fmt.Println(multiply(5, 10))
	func(x, y int) {
		fmt.Println(x, y)
	}(5, 10)
	test := func(x, y int) {}
	test(5, 10)
}
