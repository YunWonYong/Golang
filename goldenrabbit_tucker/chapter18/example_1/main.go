package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice1 := []int{}
	sliceInfoPrint[int](slice1)
	slice2 := []int{ 0, 10, 30, 12 }
	sliceInfoPrint[int](slice2)
	slice3 := []int{ 0: 12, 4: 52, 10: 99 }
	sliceInfoPrint[int](slice3)
}


func sliceInfoPrint[T interface{}](slice []T) {
	sliceType := reflect.ValueOf(slice)
	fmt.Printf("slice: %#+v, len: %d, cap: %d\n", sliceType, len(slice), cap(slice))
}
