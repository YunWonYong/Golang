package main

import (
	"fmt"
	"unsafe"
)

type Data struct {
	value int
	data  [200]int
}

func main() {
	arg := &Data{
		value: 500,
	}

	arg.data[100] = 400

	fmt.Printf("Data Initial szie: %d\n", unsafe.Sizeof(arg))
	fmt.Printf("Data Initial value: %d, %d\n", arg.value, arg.data[100])

	callByValueChange(*arg)
	fmt.Printf("callByValueChange function execute after Data value: %d, %d\n", arg.value, arg.data[100])

	callByReferenceChange(arg)
	fmt.Printf("callByReferenceChange function execute after Data value: %d, %d\n", arg.value, arg.data[100])
}

func callByValueChange(arg Data) {
	fmt.Printf("callByValueChange function argument size: %d\n", unsafe.Sizeof(arg))
	arg.value = 900
	arg.data[100] = 900
}

func callByReferenceChange(arg *Data) {
	fmt.Printf("callByReferenceChange function argument size: %d\n", unsafe.Sizeof(arg))
	arg.value = 900
	arg.data[100] = 900
}
