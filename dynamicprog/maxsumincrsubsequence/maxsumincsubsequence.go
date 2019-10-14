package maxsumincrsubsequence

import "math"

func MaxSumIncreasingSubsequence(array []int) []interface{} {
	// Write your code here.
	sums := make([]int, len(array))
	sequences := make([]int, len(array))
	maxSumsIdx := 0
	for i := range array {
		sequences[i] = math.MinInt32
		sums[i] = array[i]
	}

	for i, currentNum := range array {
		for j := 0; j < i; j++ {
			otherNum := array[j]
			if otherNum < currentNum && sums[j]+currentNum >= sums[i] {
				sums[i] = sums[j] + currentNum
				sequences[i] = j
			}
		}
		if sums[i] > sums[maxSumsIdx] {
			maxSumsIdx = i
		}
	}
	sequence := buildSequence(array, sequences, maxSumsIdx)

	return []interface{}{sums[maxSumsIdx], sequence}
}

func buildSequence(array []int, sequences []int, currentIdx int) []int {
	sequence := []int{}
	for currentIdx != math.MinInt32 {
		sequence = append(sequence, array[currentIdx])
		currentIdx = sequences[currentIdx]
	}
	reverse(sequence)
	return sequence
}

func reverse(array []int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}
