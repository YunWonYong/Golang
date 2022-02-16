package main

import "fmt"

func main() {
	a := make([]int, 5)
	b := make([]int, 5, 10)
	fmt.Printf("a slice variable value: %v, tpye: %T, length: %d, capacity: %d\n", a, a, len(a), cap(a))
	fmt.Printf("b slice variable value: %v, tpye: %T, length: %d, capacity: %d\n", b, b, len(b), cap(b))
	c := []int{}
	d := []int{1, 2, 3, 4, 5}
	fmt.Printf("c slice variable value: %v, tpye: %T, length: %d, capacity: %d\n", c, c, len(c), cap(c))
	fmt.Printf("d slice variable value: %v, tpye: %T, length: %d, capacity: %d\n", d, d, len(d), cap(d))

	e := make([]int, 4, 5)
	e[0] = 1
	e[1] = 2
	e[2] = 3
	e[3] = 4
	fmt.Printf("e slice variable value: %v, tpye: %T, length: %d, capacity: %d\n", e, e, len(e), cap(e))
	f := make([]int, 4, 5)
	f[0] = 1
	f[1] = 2
	f[2] = 3
	f[3] = 4
	//f[4] = 5 //error
	fmt.Printf("f slice variable value: %v, tpye: %T, length: %d, capacity: %d\n", f, f, len(f), cap(f))
}
