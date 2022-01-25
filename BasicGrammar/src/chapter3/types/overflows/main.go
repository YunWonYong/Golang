package main

import (
	"fmt"
	"math"
)

func main() {
	OverflowCasting()
}

func OverflowCasting() {
	var a int16 = math.MaxInt8 + 1
	var b int8 = int8(a)
	fmt.Printf("a: %d, b: %d\n", a, b)
}
