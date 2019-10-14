package main

import (
	"fmt"

	"github.com/maneeshchaturvedi/gotheotherway/ds-algo/array/longestrange"
)

func main() {
	arr := []int{4, 2, 3, 1, 6, 7, 8, 9, 10, 11}
	fmt.Println(longestrange.LargestRange(arr))
}
