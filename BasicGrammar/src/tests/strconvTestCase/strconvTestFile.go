package strconvTest

import "strconv"

func StringToBoolean(str string) bool {
	result, _ := strconv.ParseBool(str)
	return result
}

func StringToInt64(str string) int64 {
	result, _ := strconv.ParseInt(str, 10, 64) // 문자열, n진수, bit표현 범위
	return result
}

func StringToInt(str string) int {
	result, _ := strconv.Atoi(str)
	return result
}
