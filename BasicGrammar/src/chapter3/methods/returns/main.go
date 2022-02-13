package main

import "fmt"

func main() {
	x := GetX()
	y := GetY()
	divisor, remain := Div(x, y)
	fmt.Println(divisor, remain)
	divisor, remain = Div2(x, y)
	fmt.Println(divisor, remain)
	divisor, remain = Div3(x, y)
	fmt.Println(divisor, remain)

	z, err := ErrorFn(x, y)
	if err == nil {
		fmt.Println("Error!!!")
	}
	fmt.Println(z)
}

func GetX() int {
	return 16
}

func GetY() (y int) {
	y = 9
	return
}

func Div(x, y int) (int, int) {
	return x % y, x / y
}

func Div2(x, y int) (int, int) {
	divisor := x % y
	remain := x / y
	return divisor, remain
}

func Div3(x, y int) (divisor, remain int) {
	divisor = x % y
	remain = x / y
	return
}

func ErrorFn(x, y int) (z int, err interface{}) {
	z = x + y
	err = nil
	return
}
