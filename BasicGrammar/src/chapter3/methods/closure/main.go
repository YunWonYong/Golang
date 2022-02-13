package main

import (
	"fmt"
	"strings"
)

func main() {
	f := InfinityLoop("수박")
	n := 10
	answer := ""
	for {
		if n == 0 {
			break
		}
		answer += f()
		n--
	}
	fmt.Println(answer)

	// f = later() // 타입이 같지 않아 대입 불가
	f2 := later()
	fmt.Print(f2("Hello"))
	fmt.Print(f2(" "))
	fmt.Print(f2("world"))
	fmt.Println(f2(" "))
}

func InfinityLoop(str string) func() string {
	arr := strings.Split(str, "")
	index, length := 0, len(arr)
	return func() (el string) {
		if index == length {
			index = 0
		}
		el = string(arr[index])
		index++
		return
	}
}

func later() func(string) string {
	store := ""
	return func(next string) (s string) {
		s = store
		store = next
		return
	}
}
