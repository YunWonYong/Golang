package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	IntegerTypeMaxAndMinValuePrint()
	ImpliciteDefaultIntType()
	IntegerExpression()
	IntegerTypeCasting()
}

func IntegerTypeMaxAndMinValuePrint() {
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
}

func ImpliciteDefaultIntType() {
	fmt.Println("===================================================")
	fmt.Printf("reflect.TypeOf(1) => %v\n", reflect.TypeOf(1))
	fmt.Println("===================================================")
}

func IntegerExpression() {
	fmt.Printf("00001111 => %d\n", 00001111) // 2진수는 안됌 앞에 0이 붙기 때문에 8진수의 결과로 나옴
	fmt.Printf("15 => %d\n", 15)
	fmt.Printf("017 => %d\n", 017)
	fmt.Printf("0xF => %d\n", 0xF)
	fmt.Printf("0xf => %d\n", 0xf)
	fmt.Println("===================================================")
}

func IntegerTypeCasting() {
	var i8 int8 = 1
	var i16 int16 = int16(i8)
	var i32 int32 = int32(i16)
	var i64 int64 = int64(i32)
	var i int = int(i8)

	fmt.Printf("i := 1 => %T\n", i) // 2진수는 안됌 앞에 0이 붙기 때문에 8진수의 결과로 나옴
	fmt.Printf("int8(i) => %T\n", int8(i))
	fmt.Printf("int16(i) => %T\n", int16(i))
	fmt.Printf("int32(i) => %T\n", int32(i))
	fmt.Printf("int64(i) => %T\n", int64(i))
	fmt.Println("===================================================")

	fmt.Println(i8, i16, i32, i64)
}
