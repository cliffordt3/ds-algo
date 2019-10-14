package maxsumsubarraysizek

import "math"

/*
find the maximum sum of a contiguous subarray of size k
assume k < len(array)
sliding window approach
*/

func maxSubArraySum(array []int, k int) int {
	currentRunningSum := 0
	maxSubArraySum := math.MinInt32
	for i := 0; i < len(array)-1; i++ {
		currentRunningSum += array[i]
		if i >= k-1 {
			maxSubArraySum = max(maxSubArraySum, currentRunningSum)
			currentRunningSum -= array[i-(k-1)] //since the array index is 0 based
		}
	}
	return maxSubArraySum
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return n
}
