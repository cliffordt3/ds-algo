package smallestsubarraywithgivensum

import "math"

/* find the smallest subarray size with given sum
[4,2,2,7,8,1,2,8,1,0] find smallest sub array with sum >=8
*/
func smallestSubArraySizeWithGivenSum(array []int, target int) int {
	currentWindowSum := 0
	minWindowSize := math.MaxInt32
	windowStart := 0
	for windowEnd := 0; windowEnd < len(array)-1; windowEnd++ {
		currentWindowSum += array[windowEnd]
		for currentWindowSum >= target {
			minWindowSize = min(minWindowSize, windowEnd-windowStart+1)
			currentWindowSum -= array[windowStart]
			windowStart++
		}
	}
	return minWindowSize
}
func min(n, m int) int {
	if n < m {
		return n
	}
	return m
}
