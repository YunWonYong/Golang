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
