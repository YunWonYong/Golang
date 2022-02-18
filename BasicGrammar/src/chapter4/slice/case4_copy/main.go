package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, 2, 3)
	copyElementCount := copy(b, a)
	if copyElementCount != 2 {
		panic("망했죠~")
	}
	fmt.Printf("a: %#v, b: %#v\n", a, b)

	//c := make([]string, 9) //error
	//c := make([]int32, 9) //error
	//c := make([]rune, 9) //error
	c := make([]byte, 9)
	str := "가나다라마바사아자차카타파하"
	copyElementCount = copy(c, str)
	if copyElementCount != 9 {
		panic("망했죠~")
	}
	fmt.Printf("c: %#v, str: %#v\n", string(c), str)

	c = make([]byte, 9)
	str = "ABCDEFGHIJKLMNOP"
	copyElementCount = copy(c, str)
	if copyElementCount != 9 {
		panic("망했죠~")
	}
	fmt.Printf("c: %#v, str: %#v\n", string(c), str)
}
