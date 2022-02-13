package main

import "fmt"

func main() {
	MyFunction(1, 2, 3, 4, 5)
	MyFunction(1, 2, 3)

	//MyFunction2(1) error
	MyFunction2(1, 2)
}

func MyFunction2(i, _ int) {
	fmt.Println(i)
}

func MyFunction(i ...int) {
	fmt.Printf("%#v\n", i)
}

/* Go는 overloading을 지원하지 않음.
func MyFunction(str string) {

}
func MyFunction(i, j int) {

}
*/
