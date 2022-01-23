package signedInteger

import (
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
