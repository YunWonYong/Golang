package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice1 := append([]string{})
	sliceInfoPrint[string](slice1)
	slice2 := append([]string{}, "A", "B", "C", "D")
	sliceInfoPrint[string](slice2)
	slice3 := append([]int64{}, 65, 66, 67, 68)
	sliceInfoPrint[int64](slice3)

	for i:= 0; i < 26; i++ {
		if i == len(slice2) {
			slice2 = append(slice2, string(65 + i))
		}
		fmt.Printf("slice len: %d, \tcap %d,\tindex: %d, \tvalue: %s,\tpointer: %p\n", len(slice2), cap(slice2), i, slice2[i], slice2)
	}
}


func sliceInfoPrint[T interface{}](slice []T) {
	sliceType := reflect.ValueOf(slice)
	fmt.Printf("slice: %#+v, len: %d, cap: %d\n", sliceType, len(slice), cap(slice))
}
