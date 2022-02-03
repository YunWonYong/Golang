package main

import "fmt"

func main() {
	ArrayDeclaration()
}

func ArrayDeclaration() {
	//var a [5]int = [5]int{1, 2, 3, 4, 5}

	var a [5]int = [...]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	var d [5]int8
	d[0] = 1
	d[1] = 2
	d[2] = 3
	d[3] = 4
	d[4] = 5

	var b = [5]int{1, 2, 3, 4, 5}
	var c = [5]int{1, 2, 3}
	var f = [...]int{1, 2}

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)

	e := [5]int{1, 2, 3, 4, 5}
	h := [...]int{1, 2, 3}
	i := [5]int{}
	fmt.Println(e)
	fmt.Println(h)
	fmt.Println(i)

	//j := []int{}
	//j[0] = 1 // error

	x := [3]int8{1, 2, 3}
	//x[0] = 128 // error

	y := [3]int16{1, 2, 3}
	//x = y // error
	//y = x // error

	fmt.Println(x, y)

	z := x
	fmt.Printf(`
 x variable momey address => %p, elements => %#v
 z variable momey address => %p, elements => %#v
 `, &x, x, &z, z)

	z[0] = 10
	z[1] = 20
	z[2] = 30
	fmt.Printf(`
 x variable momey address => %p, elements => %#v
 z variable momey address => %p, elements => %#v
 `, &x, x, &z, z)
}
