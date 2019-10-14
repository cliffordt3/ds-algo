package longestrange

import "sort"

func LargestRange(array []int) []int {
	nums := map[int]bool{}
	longest := 0
	best := []int{}
	for _, num := range array {
		nums[num] = true
	}

	for _, num := range array {
		if !nums[num] {
			continue
		}
		nums[num] = false
		currentLen := 1
		left := num - 1
		right := num + 1
		for nums[left] {
			nums[left] = false
			currentLen++
			left--
		}
		for nums[right] {
			nums[right] = false
			currentLen++
			right++
		}
		if currentLen > longest {
			longest = currentLen
			best = []int{left + 1, right - 1}
		}
	}
	return best
}

