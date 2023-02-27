package example

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/integralist/go-findroot/find"
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
	fmt.Println(os.Executable())

	root, err := find.Repo()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	fmt.Printf("%+v", root)

	return filepath.Abs(".")
}
