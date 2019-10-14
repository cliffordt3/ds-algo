package twonumsum

import "sort"

func TwoNumberSum(array []int, target int) []int {
	for i := 0; i < len(array)-1; i++ {
		first := array[i]
		for j := i + 1; j < len(array); j++ {
			second := array[j]
			if first+second == target {
				results := []int{first, second}
				sort.Ints(results)
				return results
			}
		}
	}
	return []int{}
}
func TwoNumberSumV2(array []int, target int) []int {
	m := make(map[int]bool)
	for _, num := range array {
		secondNum := target - num
		if _, ok := m[secondNum]; ok {
			results := []int{num, secondNum}
			sort.Ints(results)
			return results
		}
		m[num] = true

	}
	return []int{}
}
