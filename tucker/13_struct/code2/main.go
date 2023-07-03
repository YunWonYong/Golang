package main

import (
	"fmt"
	"unsafe"
)

type (
	SizeCheckStruct struct {
		a int32
		b int32
	}

	MemoryAlignmentCheckStruct struct {
		a int32
		b float64
	}

	MemoryAlignmentCheckStruct2 struct {
		a float64
		b int32
	}

	MemoeyPaddingCheckStruct struct {
		a int8
		b int
		c int8
		d int
		e int8
	}

	MemoeyPaddingCheckStruct2 struct {
		a int8
		c int8
		e int8
		b int
		d int
	}
)

func main() {
	scs := SizeCheckStruct{1, 2}

	fmt.Println("SizeCheckStruct: ", unsafe.Sizeof(scs))

	macs := MemoryAlignmentCheckStruct{1, 1.2}
	fmt.Println("MemoryAlignmentCheckStruct: ", unsafe.Sizeof(macs)) // memory padding

	macs2 := MemoryAlignmentCheckStruct2{1.2, 1}
	fmt.Println("MemoryAlignmentCheckStruct2: ", unsafe.Sizeof(macs2)) // memory padding

	mpcs := MemoeyPaddingCheckStruct{1, 2, 3, 4, 5}
	fmt.Println("MemoeyPaddingCheckStruct: ", unsafe.Sizeof(mpcs)) // memory padding

	mpcs2 := MemoeyPaddingCheckStruct2{1, 2, 3, 4, 5}
	fmt.Println("MemoeyPaddingCheckStruct2: ", unsafe.Sizeof(mpcs2))
}
