package main

import "fmt"

func main() {
	defer RecoverWrapFn(func() { fmt.Println("잘가~") })
	//panic(nil) // nil 넣으면 종료해버림
	panic("갈땐 가더라도 Print함수 호출해줘!!")
}

func RecoverWrapFn(callback func()) {
	x := recover()
	if x != nil {
		fmt.Println(x)
		callback()
		return
	}
	fmt.Printf("x: %#v", x)
}
