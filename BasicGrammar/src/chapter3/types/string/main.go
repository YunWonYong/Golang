package main

import (
	"fmt"
	"math"
)

func main() {
	StringTypeInitValues()
	fmt.Println("============================")
	fmt.Println("\101\x41\u0041\U00000041")
	fmt.Println(`\101\x41\u0041\U00000041`)

	fmt.Printf(`
int8  MaxValue: %d
int16 MaxValue: %d
int32 MaxValue: %d
int64 MaxValue: %d
`, math.MaxInt8, math.MaxInt16, math.MaxInt32, math.MaxInt64)

	s := "test"
	fmt.Println(s[0])

	//s[0] = "T"
	//fmt.Println(s[0])
}

func StringTypeInitValues() {
	var s1 string = string([]rune{'a', 'b', 'c'})
	fmt.Println(s1)
	s2 := "hello"
	fmt.Println(s2)
	s3 := `hello`
	fmt.Println(s3)
}
