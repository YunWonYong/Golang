package unsignedInteger

import (
	"testing"
)

func TestByteOrUint8(t *testing.T) {
	if ByteOrUint8() == "byte" {
		t.Error("no uint8 type")
	}
}
