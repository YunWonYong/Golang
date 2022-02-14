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
	fmt.Println("-----------------------")
	ConstDefaultValueOmit()
	fmt.Println("-----------------------")
	ConstInitialValueFormula()
}

func ConstDefaultValueOmit() {
	const (
		X   = 1
		Y   // ?
		Z   // ?
		STR = "가"
		STR2
		STR3 = "나"
		STR4
	)
	fmt.Println("number default value omit: ", X, Y, Z)
	fmt.Println("string default value omit: ", STR, STR2, STR3, STR4)
}

func ConstInitialValueFormula() {
	const (
		X = 5
		Y = 10
		Z = X + Y
	)
	fmt.Printf("%d + %d = %d", X, Y, Z)
	/*error!!!!!!!!!!!
	const (
		C = A + B
		A = 10
		B = 5
	)
	fmt.Printf("%d + %d = %d", A, B, C)
	*/
}
