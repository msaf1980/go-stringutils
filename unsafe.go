package stringutils

import (
	"reflect"
	"unsafe"
)

// UnsafeString returns the string with specific length under byte buffer.
func UnsafeString(ptr *byte, length int) (s string) {
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	str.Data = uintptr(unsafe.Pointer(ptr))
	str.Len = length

	return s
}
