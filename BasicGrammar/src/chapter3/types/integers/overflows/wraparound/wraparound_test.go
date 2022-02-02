package wraparound

import (
	"math"
	"testing"
)

func TestIntAround(t *testing.T) {
	actual := IntAround()
	if actual != math.MinInt8 {
		t.Errorf("actual: %d, type: %T", actual, actual)
	}
}

func TestIntAroundLoop(t *testing.T) {
	var (
		except, actual, index, length int8 = math.MinInt8, 0, 0, math.MaxInt8
	)

	for index < length {
		index++
		actual = IntAroundLoop(index)
		if except != actual {
			t.Errorf("actual: %d, except: %d", actual, except)
		}
		except++
	}
}
