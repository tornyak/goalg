package util

import (
	"runtime"
	"reflect"
)

// GetFuncName returns the name of the calling function
func GetCallingFuncName() string {
	pc := make([]uintptr, 10) // at least 2 entries needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func GetFuncName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
