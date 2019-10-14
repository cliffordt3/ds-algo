package subarraysort

import (
	"math"
)

func SubarraySort(array []int) []int {
	minOutOfOrder := math.MaxInt32
	maxOutOfOrder := math.MinInt32
	for i, num := range array {
		if isOutOfOrder(i, num, array) {
			minOutOfOrder = min(minOutOfOrder, num)
			maxOutOfOrder = max(maxOutOfOrder, num)
		}
	}
	if minOutOfOrder == math.MaxInt32 {
		return []int{-1, -1}
	}
	subArrayStartIdx := 0
	for minOutOfOrder >= array[subArrayStartIdx] {
		subArrayStartIdx++
	}
	subArrayEndIdx := len(array) - 1
	for maxOutOfOrder <= array[subArrayEndIdx] {
		subArrayEndIdx--
	}
	return []int{subArrayStartIdx, subArrayEndIdx}
}

func isOutOfOrder(i int, num int, array []int) bool {
	if i == 0 {
		return num > array[i+1]
	}
	if i == len(array)-1 {
		return num < array[i-1]
	}
	return num > array[i+1] || num < array[i-1]
}

func min(minOutOfOrder int, num int) int {
	if minOutOfOrder < num {
		return minOutOfOrder
	}
	return num
}

func max(maxOutOfOrder int, num int) int {
	if maxOutOfOrder > num {
		return maxOutOfOrder
	}
	return num
}
