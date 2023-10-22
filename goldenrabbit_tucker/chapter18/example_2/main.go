package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice1 := make([]string, 0)
	sliceInfoPrint[string](slice1)
	slice2 := make([]string, 5, 10)
	sliceInfoPrint[string](slice2)
	slice3 := make([]string, 0, 100)
	sliceInfoPrint[string](slice3)

	for i:= 0; i < len(slice2); i++ {
		slice2[i] = string(65 + i)
		fmt.Printf("index: %d, value: %s\n", i, slice2[i])
	}
}


func sliceInfoPrint[T interface{}](slice []T) {
	sliceType := reflect.ValueOf(slice)
	fmt.Printf("slice: %#+v, len: %d, cap: %d\n", sliceType, len(slice), cap(slice))
}
