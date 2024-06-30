package main

import (
	"fmt"
)

func main() {

	var s = "gfcaaaecbg"
	var n int64 = 547602

	var aCount float64
	var aCountInS int64
	for _, v := range s {
		if v == 'a' {
			aCountInS++
		}
	}
	aCount = float64((n * aCountInS)) / float64(len(s))
	fmt.Println(aCount)
}
