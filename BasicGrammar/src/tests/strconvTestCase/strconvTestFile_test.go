package strconvTest

import "testing"

func TestStringToBoolean(t *testing.T) {
	expect := true
	actual := StringToBoolean("true")

	if expect != actual {
		t.Errorf("%t != %t", expect, actual)
	}
}

func TestStringToInt64(t *testing.T) {
	var expect int64 = 9999
	actual := StringToInt64("9999")

	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
	}
}

func TestStringToInt(t *testing.T) {
	expect := 10000
	actual := StringToInt("10000")

	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
	}
}
