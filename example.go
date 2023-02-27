package example

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func DirPath() string {
	_, b, _, _ := runtime.Caller(0)
	curDir := filepath.Dir(b)
	return curDir
}

func WDPath() (string, error) {
	fmt.Println(filepath.Abs("."))
	fmt.Println(path.Dir("."))
	fmt.Println(os.Getwd())
	return filepath.Abs(".")
}
