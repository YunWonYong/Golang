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
	var a int
	fmt.Println("선언된 a라는 int 타입의 변수값: ", a)
}

func variableInitialized() {
	var a int = 10
	fmt.Printf("선언된 a라는 int 타입의 변수에 %d라는 값으로 초기화\n", a)
}

func multiVariableDefinition() {
	var (
		a int = 100
		b int = 200
		c float64 
	)

	fmt.Printf("a: %d, b: %d, c: %f\n", a, b, c)
}
