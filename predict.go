// Package setu setu-class-cpp 的 go 接口
package setu

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/lib -lsetu
// #include <stdlib.h>
// #include <setu.h>
import "C"
import (
	"unsafe"
)

func LoadModule(path string) int {
	cpth := C.CString(path)
	defer C.free(unsafe.Pointer(cpth))
	return (int)(C.load_module(cpth))
}

func PredictFile(path string, index int) int {
	cpth := C.CString(path)
	defer C.free(unsafe.Pointer(cpth))
	return (int)(C.predict_file(cpth, (C.int)(index)))
}
