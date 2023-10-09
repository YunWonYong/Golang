package main

import (
	"fmt"
	"strings"
)

func init() {
	fmt.Println("init 함수는 main 함수보다 먼저 호출됨")
}

func main() {
	fmt.Println("main 함수 호출됨")
	voidFn()
	paramFn(10, "hello world!!!")
	argumentsFn(20, "h", "ello", " wor", "l", "d", "!!!")
	fmt.Printf("singleReturnFn call!!! return value: %#v\n", singleReturnFn())


	a, b := multiReturnFn()
	fmt.Printf("multiReturnFn call!! return firstValue: %d, secondValue: %s\n", a, b)
	a, b = multiReturnFn2()
	fmt.Printf("multiReturnFn2 call!! return firstValue: %d, secondValue: %s\n", a, b)
	a, b = multiReturnFn3()
	fmt.Printf("multiReturnFn3 call!! return firstValue: %d, secondValue: %s\n", a, b)

	a, _ = multiReturnFn2()
	fmt.Printf("multiReturnFn2 call!! return firstValue: %d, secondValue: %s\n", a, b)

	a, b = all(100, "100")
	fmt.Printf("all call!! firstValue: %d, secondValue: %s\n", a, b)
}

func voidFn() {
	fmt.Println("voidFn call!!! 아무것도 받지 않고 아무것도 반환하지 않음")
}

func paramFn(a int64, s string) {
	fmt.Printf("paramFn call!!! firstParam: %d, secondParam: %s\n", a, s)
}

func argumentsFn(a int64, s ...string) {
	fmt.Printf("argumentsFn call!! firstParam: %d, params: %#v, join: %s\n", a, s, strings.Join(s, ""))
}

func singleReturnFn() int {
	return 10
}

func multiReturnFn() (int, string) {
	return 20, "20"
}

func multiReturnFn2() (a int, b string) {
	a, b = multiReturnFn()
	return 
}

func multiReturnFn3() (a int, b string) {
	a, b = multiReturnFn()
	return 30, "30"
}

func all(a int, b string) (int, string) {
	return a, b
}
