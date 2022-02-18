package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sumValue := sum(a...)
	fmt.Printf("a %#v element sum: %d\n", a, sumValue)
	a = append(a, 10)
	sumValue = sum(a...)
	fmt.Printf("a %#v element sum: %d\n", a, sumValue)
}

func sum(a ...int) (result int) {
	for _, element := range a {
		result += element
	}
	return
}
