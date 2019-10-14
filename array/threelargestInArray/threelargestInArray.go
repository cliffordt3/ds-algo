package threelargestInArray

import (
	"math"
)

func FindThreeLargestNumbers(array []int) []int {
	threeLargest := []int{math.MinInt32, math.MinInt32, math.MinInt32}
	for _, num := range array {
		updateLargest(threeLargest, num)
	}
	return threeLargest
}

func updateLargest(threeLargest []int, num int) {
	if num > threeLargest[2] {
		shiftAndUpdate(threeLargest, num, 2)
	} else if num > threeLargest[1] {
		shiftAndUpdate(threeLargest, num, 1)
	} else if num > threeLargest[0] {
		shiftAndUpdate(threeLargest, num, 0)
	}
}

func shiftAndUpdate(array []int, num int, idx int) {
	for i := 0; i < idx+1; i++ {
		if i == idx {
			array[i] = num
		} else {
			array[i] = array[i+1]
		}
	}
}
