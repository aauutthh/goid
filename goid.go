package goid

import (
	"reflect"
	"unsafe"
)

func getg() interface{}
func Getg() uintptr
func GetGoIDAsm() int64

var goidOffset uintptr = 0

func init() {
	f, ok := reflect.TypeOf(getg()).FieldByName("goid")
	if ok {
		goidOffset = f.Offset
	}
}

func GetGoID() int64 {
	if goidOffset > 0 {
		goid := *(*int64)(unsafe.Pointer(Getg() + goidOffset))
		return goid
	}
	return -1
}
