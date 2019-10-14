package smallestdiff

import (
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	counter1, counter2 := 0, 0
	current, smallest := math.MaxInt32, math.MaxInt32
	result := []int{}
	sort.Ints(array1)
	sort.Ints(array2)
	for counter1 < len(array1) && counter2 < len(array2) {
		first, second := array1[counter1], array2[counter2]
		if first < second {
			current = second - first
			counter1++
		} else if second < first {
			current = first - second
			counter2++
		} else {
			return []int{first, second}
		}
		if current < smallest {
			smallest = current
			result = []int{first, second}
		}
	}
	return result
}
