package unsignedInteger

import (
	"reflect"
)

func ByteOrUint8() string {
	i := byte(1)
	return reflect.TypeOf(i).Name()
}
