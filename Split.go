package main

import (
	"strings"
	"fmt"
)

func SplitLogs(data []byte) {
	Log := string(data)
	a := strings.Split(Log, "\n")
	fmt.Println(a)
}
