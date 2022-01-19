package main

import "fmt"

func main() {
	var var1 int = 1
	fmt.Println(var1)
	var (
		var2, var3, var4           int    = 2, 3, 3
		name, email, phone, gender string = "원용", "ywyi1992@gmail.com", "112", "남"
	)
	fmt.Printf("정수 타입의 변수: %#v, %#v, %#v, %#v\n", var1, var2, var3, var4)  // 타입에 상관없이 값을 보여줌
	fmt.Printf("문자열 타입의 변수: %T, %T, %T, %T\n", name, email, phone, gender) // 변수의 타입을 출력함
}
