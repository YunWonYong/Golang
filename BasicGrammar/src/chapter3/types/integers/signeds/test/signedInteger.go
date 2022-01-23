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
