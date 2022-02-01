package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	OverflowCasting()
}

func OverflowCasting() {
	var a int16 = math.MaxInt8
	fmt.Printf("in8Max: %d, binaryCode: %08s\n", a, strconv.FormatInt(int64(a), 2))
	fmt.Printf("in8Max + 1: %d, binaryCode: %09s\n", a, strconv.FormatInt(int64(a + 1), 2))
	var b int8 = int8(a + 1)
	fmt.Printf("wrap around: %d, b: %09s\n", b, strconv.FormatInt(int64(b), 2))
}
