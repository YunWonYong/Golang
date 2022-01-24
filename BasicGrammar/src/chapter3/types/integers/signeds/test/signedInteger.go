package signedInteger

import (
	"math"
	"reflect"
)

func IntDefaultType() string {
	i := 1
	return reflect.TypeOf(i).Name()
}
func IntOfInt64() bool {
	var intVariable int = 0
	var int64Variable int64 = 0
	return reflect.TypeOf(intVariable) == reflect.TypeOf(int64Variable)
}

func IntOrInt64() string {
	i := math.MaxInt64
	return reflect.TypeOf(i).Name()
}

func IntAround() (result int8) {
	result = math.MaxInt8
	result++
	return
}

func IntAroundLoop(sumValue int8) int8 {
	return math.MaxInt8 + sumValue
}
