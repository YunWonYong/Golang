package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var int8Min int8 = math.MinInt8
	var int16Min int16 = math.MinInt16
	var int32Min int32 = math.MinInt32
	var int64Min int64 = math.MinInt64
	fmt.Printf("int8 Min Value: %d, type: %T\n", int8Min, int8Min)
	fmt.Printf("int16 Min Value: %d, type: %T\n", int16Min, int16Min)
	fmt.Printf("int32 Min Value: %d, type: %T\n", int32Min, int32Min)
	fmt.Printf("int64 Min Value: %d, type: %T\n", int64Min, int64Min)
	fmt.Println("===================================================")
	var int8Max int8 = math.MaxInt8
	var int16Max int16 = math.MaxInt16
	var int32Max int32 = math.MaxInt32
	var int64Max int64 = math.MaxInt64
	fmt.Printf("int8 Min Value: %d, type: %T\n", int8Max, int8Max)
	fmt.Printf("int16 Min Value: %d, type: %T\n", int16Max, int16Max)
	fmt.Printf("int32 Min Value: %d, type: %T\n", int32Max, int32Max)
	fmt.Printf("int64 Min Value: %d, type: %T\n", int64Max, int64Max)
	fmt.Println("===================================================")
	impliciteIntMin := math.MinInt
	impliciteIntMax := math.MaxInt
	fmt.Printf("implicite int Min Value: %d, type: %T\n", impliciteIntMin, impliciteIntMin)
	fmt.Printf("implicite int Max Value: %d, type: %T\n", impliciteIntMax, impliciteIntMax)
	var explicateIntMin int = math.MinInt
	var expliciteIntMax int = math.MaxInt
	fmt.Printf("explicite int Min Value: %d, type: %T\n", explicateIntMin, explicateIntMin)
	fmt.Printf("explicite int Max Value: %d, type: %T\n", expliciteIntMax, expliciteIntMax)

	fmt.Println(reflect.TypeOf(1))
}
