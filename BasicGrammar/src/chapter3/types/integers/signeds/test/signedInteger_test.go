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
