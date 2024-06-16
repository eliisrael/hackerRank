package main

import (
	"fmt"
	"strings"
)

func main() {

	var n int32 = 6
	stairCase(n)
}

func stairCase(n int32) {
	var i int32
	for i = 0; i < n; i++ {
		fmt.Println(strings.Repeat(" ", int(n-i-1)) + strings.Repeat("#", int(i+1)))
	}
}
