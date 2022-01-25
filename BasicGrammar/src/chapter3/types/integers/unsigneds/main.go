package main

import (
	"fmt"
	"math"
)

func main() {
	UintegerTypeMaxValuePrint()
	UinutegerTypeCasting()
}

func UintegerTypeMaxValuePrint() {
	var uint8Max uint8 = math.MaxUint8
	var uint16Max uint16 = math.MaxInt16
	var uint32Max uint32 = math.MaxInt32
	var uint64Max uint64 = math.MaxInt64
	var uintMax uint = math.MaxUint
	fmt.Printf("uint8 Max Value: %d, type: %T\n", uint8Max, uint8Max)
	fmt.Printf("uint16 Max Value: %d, type: %T\n", uint16Max, uint16Max)
	fmt.Printf("uint32 Max Value: %d, type: %T\n", uint32Max, uint32Max)
	fmt.Printf("uint64 Max Value: %d, type: %T\n", uint64Max, uint64Max)
	fmt.Printf("uintMax Max Value: %d, type: %T\n", uintMax, uintMax)
	fmt.Println("===================================================")
}

func UinutegerTypeCasting() {
	var ui8 uint8 = 1
	var ui16 uint16 = uint16(ui8)
	var ui32 uint32 = uint32(ui16)
	var ui64 uint64 = uint64(ui32)
	i := uint(ui8)
	var b8 byte = byte(ui8)
	fmt.Printf("var ui8 uint8 = 1 => %T\n", ui8)
	fmt.Printf("uint16(ui8) => %T\n", ui16)
	fmt.Printf("uint32(ui16) => %T\n", ui32)
	fmt.Printf("uint64(ui32) => %T\n", ui64)
	fmt.Printf("uint(ui8) => %T\n", i)
	fmt.Printf("byte(ui8) => %T\n", b8)
	b8 = ui8
	fmt.Printf("b8 = ui8 => %T\n", b8)
	fmt.Printf("byte(b8) => %T\n", byte(b8))
	fmt.Println("===================================================")
}
