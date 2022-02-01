package main

import (
	"fmt"
	"math"
)

func main() {
	//var realNumber float = 0.0
	FloatsVariable()
	fmt.Println("================================================")
	MaxAndMinValuePrint()
	fmt.Println("================================================")
	NaNAndInfinityPrint()
	Float32PrecisionCheck()
	Float64PrecisionCheck()
}

func FloatsVariable() {
	f1 := 10.0
	fmt.Printf("implicitly declared variable f1 value: %g, type: %T\n", f1, f1)
	var f2 float32 = 10.0001
	fmt.Printf("explicitly declared variable f2 value: %g, type: %T\n", f2, f2)
}

func MaxAndMinValuePrint() {
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.SmallestNonzeroFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.SmallestNonzeroFloat64)
}

func NaNAndInfinityPrint() {
	f1 := 0.0
	f2 := 0.0
	fmt.Println(f1 / f2)
	f1 = 1.0
	fmt.Println(f1 / f2)
	f1 = -1.0
	fmt.Println(f1 / f2)
}

func Float32PrecisionCheck() {
	fmt.Println("=========== float32 Preccision ===========")
	fmt.Printf("1.0000000000000000 => %v \n", float32(1.0000000000000000))
	fmt.Printf("1.0000000000000001 => %v \n", float32(1.0000000000000001))
	fmt.Printf("1.0000000000000002 => %v \n", float32(1.0000000000000002))
	fmt.Printf("1.0000000000000003 => %v \n", float32(1.0000000000000003))
	fmt.Printf("1.0000000000000004 => %v \n", float32(1.0000000000000004))
	fmt.Printf("1.0000000000000005 => %v \n", float32(1.0000000000000005))
	fmt.Printf("1.0000000000000006 => %v \n", float32(1.0000000000000006))
	fmt.Printf("1.0000000000000007 => %v \n", float32(1.0000000000000007))
	fmt.Printf("1.0000000000000008 => %v \n", float32(1.0000000000000008))
	fmt.Printf("1.0000000000000009 => %v \n", float32(1.0000000000000009))
}

func Float64PrecisionCheck() {
	fmt.Println("=========== float64 Preccision ===========")
	fmt.Printf("1.0000000000000000 => %v \n", 1.0000000000000000)
	fmt.Printf("1.0000000000000001 => %v \n", 1.0000000000000001)
	fmt.Printf("1.0000000000000002 => %v \n", 1.0000000000000002)
	fmt.Printf("1.0000000000000003 => %v \n", 1.0000000000000003)
	fmt.Printf("1.0000000000000004 => %v \n", 1.0000000000000004)
	fmt.Printf("1.0000000000000005 => %v \n", 1.0000000000000005)
	fmt.Printf("1.0000000000000006 => %v \n", 1.0000000000000006)
	fmt.Printf("1.0000000000000007 => %v \n", 1.0000000000000007)
	fmt.Printf("1.0000000000000008 => %v \n", 1.0000000000000008)
	fmt.Printf("1.0000000000000009 => %v \n", 1.0000000000000009)
}
