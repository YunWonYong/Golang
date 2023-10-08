package main

import (
	"fmt"
)

func main() {
	// 변수 선언
	variableDefinition()

	// 변수 선언 및 초기화
	variableInitialized()

	// 여러 변수를 한번에 선언 및 초기화
	multiVariableDefinition()
}

func variableDefinition() {
	a := 0
	fmt.Println("선언된 a라는 int 타입의 변수값: ", a)
}

func variableInitialized() {
	a := 10
	fmt.Printf("선언된 a라는 int 타입의 변수에 %d라는 값으로 초기화\n", a)
}

func multiVariableDefinition() {
	a := 100
	b := 200
	c := 0.0000
	fmt.Printf("a: %d, b: %d, c: %f\n", a, b, c)
}
