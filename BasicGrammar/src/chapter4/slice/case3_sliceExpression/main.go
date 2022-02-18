package main

import "fmt"

func main() {

	fmt.Println("=====SimpleSliceExpression()=====")
	SimpleSliceExpression()
	fmt.Println("=====stringSimpleSliceExpression()=====")
	stringSimpleSliceExpression()
	fmt.Println("=====fullSliceExpression()=====")
	fullSliceExpression()
}

func SimpleSliceExpression() {
	a := []int{1, 2, 3}
	b := a[0:]
	fmt.Printf("a variable[%p] equals b variable[%p] = %t \n", &a, &b, &a == &b)
	for index, element := range a {
		equalsPrint(element, b[index])
	}
}

func stringSimpleSliceExpression() {
	str := "ABC"
	temp := str[:2]
	equalsPrint(temp, "AB")
	temp = str[1:2]
	equalsPrint(temp, "B")
	temp = str[3:]
	equalsPrint(temp, "")
	temp = str[0:]
	equalsPrint(temp, "ABC")
	str = "가나다라마바"
	temp = str[0:]
	equalsPrint(temp, "가나다라마바")
	temp = str[3:]
	equalsPrint(temp, "나다라마바")
	temp = str[9:]
	equalsPrint(temp, "라마바")
	temp = str[:9]
	equalsPrint(temp, "가나다")

}

func fullSliceExpression() {
	slice := []int{2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := slice[0:len(slice):cap(slice)]
	equalsPrint(cap(slice), cap(slice2))
	slice3 := slice[1:3]
	slice4 := slice[1:3:4]
	equalsPrint(cap(slice3), cap(slice4))
	slice3 = slice[1:3:8]
	equalsPrint(cap(slice3), 8)
}

func equalsPrint(left, right interface{}) {
	fmt.Printf("left[%#v] equals right[%#v] => %t\n", left, right, left == right)
}
