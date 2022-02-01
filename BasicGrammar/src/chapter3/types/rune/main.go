package main

import (
	"fmt"
	"strconv"
)

func main() {
	RuneTypeCheck()
	fmt.Println("===============================")
	FunctionSignatureTest()
	RuneTypeValuePrefixsPrint()
}

func RuneTypeCheck() {
	var r1 rune = 0x0121
	fmt.Printf("value: %#v, %d, type: %T\n", r1, r1, r1)
	r2 := 'A'
	fmt.Printf("value: %#v, %d, type: %T\n", r2, r2, r2)
	var r3 int32 = 'A'
	fmt.Printf("value: %#v, %d, type: %T\n", r3, r3, r3)
}

func FunctionSignatureTest() {
	fmt.Println(strconv.QuoteRune(int32(32)))
}

func RuneTypeValuePrefixsPrint() {
	fmt.Println("========= rune type print =========")
	r := 'A'
	fmt.Println(r)
	fmt.Println('A')
	fmt.Println("========= string type casting =========")
	fmt.Println(string(r))
	fmt.Println(string('\101'))
	fmt.Println(string('\x41'))
	fmt.Println(string('\u0041'))
	fmt.Println(string('\U00000041'))
	fmt.Println("========= strconv.QuoteRune 함수 사용 =========")
	fmt.Printf("65: %v\n", strconv.QuoteRune(65))
	fmt.Printf("\\101: %v\n", strconv.QuoteRune('\101'))
	fmt.Printf("\\x41: %v\n", strconv.QuoteRune('\x41'))
	fmt.Printf("\\u0041: %v\n", strconv.QuoteRune('\u0041'))
	fmt.Printf("\\U00000041: %v\n", strconv.QuoteRune('\U00000041'))
}
