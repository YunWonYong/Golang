package main

import (
	f "fmt"  //fmt 패키지의 별칭으로 f라는 명을 사용함
	. "math" //math 패키지는 아무런 참조를 안해도 접근할 수 있음
)

func main() {
	f.Println("hello")
	f.Println(Pi) // meth.Pi와 같다.
}
