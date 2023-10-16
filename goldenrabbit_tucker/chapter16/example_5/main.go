package main

import (
	"fmt"

	"github.com/YunWonYong/Golang/datastruct/linkedList"
)

func main() {
	ll := linkedList.New[int64](100, 200, 300)

	for {
		value, err := ll.Pop()
		if err != nil {
			fmt.Printf("%#+v\n", err)
			break
		}

		fmt.Println(value)

		if ll.IsEmpty() {
			break
		}
	}
}
