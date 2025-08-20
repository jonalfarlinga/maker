package common

import (
	"fmt"
)

var DebugFiles = map[string]bool{}

func DebugPrintln(filename string, a ...interface{}) {
	if DebugFiles["all"] || DebugFiles[filename] {
		fmt.Println(a...)
	}
}
