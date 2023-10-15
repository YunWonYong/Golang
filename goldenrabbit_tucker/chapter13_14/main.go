package main

import (
	"fmt"
	linkedlist "gh/chap13_14/ex1/ex/datastruct/linkedList"
)

func main() {
	ll := linkedlist.New[string]("hello", " world", "!!!")

	for {
		value, err := ll.Pop()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			break
		}

		fmt.Printf("Pop %s\n", value)

		if ll.IsEmpty() {
			break
		}
	}

	ll.PushAll("hello", " world", "!!!")

	for {
		value, err := ll.Unshift()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			break
		}

		fmt.Printf("Unshift %s\n", value)

		if ll.IsEmpty() {
			break
		}
	}

}
