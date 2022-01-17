package main

import (
	"fmt"
	_ "math" // 사용하지 않지만 _를 이용하여 포함시킬 순 있음
	"strconv"
)

func main() {
	fmt.Print(strconv.FormatInt(123, 2))
}
