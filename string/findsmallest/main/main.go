package main

import (
	"fmt"

	"github.com/maneeshchaturvedi/gotheotherway/ds-algo/string/findsmallest"
)

func main() {
	s := findsmallest.FindSmallestSubString("acbdnracbn", "abn")
	fmt.Println(s)
	s = findsmallest.FindSmallestSubString("adobecodebanc", "abc")
	fmt.Println(s)
}
