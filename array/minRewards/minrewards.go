package minRewards

func MinRewards(scores []int) int {
	// Write your code here.
	rewards := make([]int, len(scores))
	for i := range scores {
		rewards[i] = 1
	}
	for i := 1; i < len(scores); i++ {
		j := i - 1
		if scores[j] < scores[i] {
			rewards[i] = rewards[j] + 1
		} else {
			for j >= 0 && scores[j] > scores[j+1] {
				rewards[j] = max(rewards[j], rewards[j+1]+1)
				j--
			}
		}
	}
	return sum(rewards)
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}

func sum(array []int) int {
	sum := 0
	for _, num := range array {
		sum = sum + num
	}
	return sum
}

func MinRewardsV2(scores []int) int {
	rewards := make([]int, len(scores))
	for i := range scores {
		rewards[i] = 1
	}
	localMinIdxs := getLocalMinIdxs(scores)
	for localMinIdx := range localMinIdxs {
		expandFromLocalMin(localMinIdx, scores, rewards)
	}
	return sum(rewards)
}

func getLocalMinIdxs(array []int) []int {
	if len(array) == 0 {
		return []int{0}
	}
	localMinIdxs := []int{}
	for i := 0; i < len(array); i++ {
		if i == 0 && array[i] < array[i+1] {
			localMinIdxs = append(localMinIdxs, i)
		}
		if i == len(array)-1 && array[i] < array[i-1] {
			localMinIdxs = append(localMinIdxs, i)
		}
		if i == 0 || i == (len(array)-1) {
			continue
		}
		if array[i] < array[i+1] && array[i] < array[i-1] {
			localMinIdxs = append(localMinIdxs, i)
		}
	}
	return localMinIdxs

}
func expandFromLocalMin(localMinIdx int, scores []int, rewards []int) {
	leftIdx := localMinIdx - 1
	for leftIdx >= 0 && scores[leftIdx] > scores[leftIdx+1] {
		rewards[leftIdx] = max(rewards[leftIdx], rewards[leftIdx+1]+1)
		leftIdx--
	}
	rightIdx := localMinIdx + 1
	for rightIdx < len(scores) && scores[rightIdx] > scores[rightIdx-1] {
		rewards[rightIdx] = rewards[rightIdx-1] + 1
		rightIdx++
	}
}

func MinRewardsV3(scores []int) int {
	rewards := make([]int, len(scores))
	for i := range scores {
		rewards[i] = 1
	}
	for i := 1; i < len(scores); i++ {
		if scores[i] > scores[i-1] {
			rewards[i] = rewards[i-1] + 1
		}
	}
	for i := len(scores) - 2; i >= 0; i++ {
		if scores[i] > scores[i+1] {
			rewards[i] = max(rewards[i], rewards[i+1]+1)
		}
	}
	return sum(rewards)
}
