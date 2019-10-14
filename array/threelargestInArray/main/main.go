package main

import (
	"fmt"

	"github.com/maneeshchaturvedi/gotheotherway/ds-algo/array/threelargestInArray"
)

func main() {
	arr := []int{3, 2, 3, 1, 4, 5, 5, 6}
	fmt.Println(threelargestInArray.FindThreeLargestNumbers(arr))
}
