package main

import (
	"fmt"

	"./foo"
)

func main() {
	foo.DoSomething()
	// foo.doSomethin() // foo파일에서만 사용할 수 있는 함수
	fmt.Println(foo.A)
	//foo.abs // foo파일에서만 사용할 수 있는 변수
	fmt.Println(foo.N)
	//foo.m // foo파일에서만 사용할 수 있는 변수
}
