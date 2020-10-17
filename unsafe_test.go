package stringutils

import (
	"reflect"
	"testing"
)

func TestUnsafeString(t *testing.T) {
	s := "test"
	bytes := []byte(s)
	if got := UnsafeString(bytes); got != s {
		t.Errorf("UnsafeString() = '%s', want '%s'", got, s)
	}
}

func TestUnsafeStringFromPtr(t *testing.T) {
	s := "test"
	bytes := []byte(s)
	if got := UnsafeStringFromPtr(&bytes[0], len(bytes)); got != s {
		t.Errorf("UnsafeStringFromPtr() = '%s', want '%s'", got, s)
	}
}

func TestUnsafeStringBytes(t *testing.T) {
	bytes := []byte("test")
	s := string(bytes)
	if got := UnsafeStringBytes(&s); !reflect.DeepEqual(got, bytes) {
		t.Errorf("UnsafeStringBytes() = '%v', want '%v'", got, bytes)
	}
}
