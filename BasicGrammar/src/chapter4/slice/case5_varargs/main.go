package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sumValue := sum(a...)
	fmt.Printf("a %#v element sum: %d\n", a, sumValue)
	b := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// sumValue = sum(b...) // ...int는 []int 타입이다.
	sumValue = sum(b[:]...)
	fmt.Printf("b %#v element sum: %d\n", b, sumValue)
}

func sum(a ...int) (result int) {
	for _, element := range a {
		result += element
	}
	return
}
