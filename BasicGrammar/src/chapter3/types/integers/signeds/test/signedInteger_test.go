package signedInteger

import (
	"math"
	"testing"
)

func TestIntDefaultType(t *testing.T) {
	if IntDefaultType() != "int" {
		t.Error("no int type")
	}
}

func TestIntOfInt64(t *testing.T) {
	if IntOfInt64() {
		t.Error("int and int64 equals")
	}
}

func TestIntOrInt64(t *testing.T) {
	if IntOrInt64() != "int" {
		t.Error("meth.MaxInt64의 타입은 int임")
	}
}

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
