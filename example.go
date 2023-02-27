package example

import (
	"path/filepath"
	"runtime"
)

func DirPath() string {
	_, b, _, _ := runtime.Caller(0)
	curDir := filepath.Dir(b)
	return curDir
}
