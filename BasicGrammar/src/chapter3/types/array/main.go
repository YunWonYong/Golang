package main

import "fmt"

func main() {
	ArrayDeclaration()
}

func ArrayDeclaration() {
	var a []int = []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	var d [5]int8
	d[0] = 1
	d[1] = 2
	d[2] = 3
	d[3] = 4
	d[4] = 5

	var b = [5]int{1, 2, 3, 4, 5}
	var c = [5]int{1, 2, 3}
	var f = [...]int{1, 2}

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)

	e := [5]int{1, 2, 3, 4, 5}
	g := []int{1, 2, 3}
	h := [...]int{1, 2, 3}
	i := [5]int{}
	j := []int{} // 의미없는 코드와 같음...
	fmt.Println(e)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)

	//j[0] = 1 error

	x := []int8{1, 2, 3}

	//x[0] = 128 error

	y := []int16{1, 2, 3}
	//x = y error
	//y = x error

	fmt.Println(x, y)

	z := x
	fmt.Printf("z변수의 0번 째 인덱스의 값을 대입하기 전 x: %v z: %v\n", x, z)
	z[0] = 10
	fmt.Printf("z변수의 0번 째 인덱스의 값을 대입한 후 x: %v z: %v\n", x, z)
}
