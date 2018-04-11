package main

import (
	"path/filepath"
	"runtime"
)

var root string

// RootDir returns the root dir this project was run from
func RootDir() string {
	return root
}

// SetupDir initializes the private root var
func SetupDir() {
	_, thisFile, _, _ := runtime.Caller(0)
	root = filepath.Dir(thisFile)
}
