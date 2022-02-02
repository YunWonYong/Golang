package wraparound

import "math"

func IntAround() (result int8) {
	result = math.MaxInt8
	result++
	return
}

func IntAroundLoop(sumValue int8) int8 {
	return math.MaxInt8 + sumValue
}
