package main

import (
	"fmt"

	"github.com/maneeshchaturvedi/gotheotherway/ds-algo/array/fournumberSum"
)

func main() {
	arr := []int{7, 6, 4, -1, 1, 2}
	target := 16

	fmt.Println(fournumberSum.FourNumberSum(arr, target))
}
