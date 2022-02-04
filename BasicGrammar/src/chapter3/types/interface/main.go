package main

import "fmt"

func main() {
	var a interface{} = 1
	var b interface{} = 1.0
	var c interface{} = '1'
	var d interface{} = "1"
	var e interface{} = [5]int{1, 1, 1, 1, 1}
	fmt.Println(a, b, c, d, e)

	a = b
	fmt.Println(a, b)

	f := e
	fmt.Println(f, e)

	fmt.Printf(`
 a variable type => %T
 b variable type => %T
 c variable type => %T
 d variable type => %T
 e variable type => %T
 f variable type => %T
 `, a, b, c, d, e, f)

	//a = b
	//fmt.Println(int(a) + int(b))
	//e[0] = 100

	fmt.Printf(`
 e array variable memory address => %p
 f array variable memory address => %p
	`, &e, &f)

	var x interface{}
	fmt.Println(x)

	x = nil
	//z := nil

}
