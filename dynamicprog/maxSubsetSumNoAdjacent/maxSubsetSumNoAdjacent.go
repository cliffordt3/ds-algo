package maxSubsetSumNoAdjacent

func MaxSubsetSumNoAdjacent(array []int) int {
	if len(array) == 0 {
		return 0
	} else if len(array) == 1 {
		return array[0]
	}
	maxSums := make([]int, len(array))

	maxSums[0] = array[0]
	maxSums[1] = max(array[0], array[1])
	for i := 2; i < len(array); i++ {
		maxSums[i] = max(maxSums[i-1], maxSums[i-2]+array[i])
	}
	return maxSums[len(maxSums)-1]
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
