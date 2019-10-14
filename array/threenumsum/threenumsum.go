package threenumsum

import "sort"

func ThreeNumberSum(array []int, target int) [][]int {

	sort.Ints(array)

	results := [][]int{}
	for i := 0; i < len(array)-2; i++ {
		second, last := i+1, len(array)-1
		for second < last {
			if array[i]+array[second]+array[last] < target {
				second++
			} else if array[i]+array[second]+array[last] > target {
				last--
			} else if array[i]+array[second]+array[last] == target {
				res := []int{array[i], array[second], array[last]}
				results = append(results, res)
				second++
				last--
			}
		}
	}
	return results
}
