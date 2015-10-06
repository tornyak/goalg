package util

import (
	"runtime"
)

// GetFuncName returns the name of the calling function
func GetFuncName() string {
	pc := make([]uintptr, 10) // at least 2 entries needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
