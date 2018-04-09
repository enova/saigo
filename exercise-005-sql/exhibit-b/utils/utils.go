package utils

import (
	"fmt"
	"runtime"
	"strings"
)

func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	}

	return original[i+1:]
}

// WhereAmI prints debug info about file, line and method name
//
func WhereAmI(depth int) string {
	function, file, line, _ := runtime.Caller(depth)
	return fmt.Sprintf("%s:%d %s()", chopPath(file), line, runtime.FuncForPC(function).Name())
}