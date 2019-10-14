package fournumberSum

func FourNumberSum(array []int, target int) [][]int {
	allPairs := map[int][][]int{}
	quadruplets := [][]int{}

	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			currentNum := array[i] + array[j]
			targetNum := target - currentNum
			if pairs, ok := allPairs[targetNum]; ok {
				for _, pair := range pairs {
					newQaudruplet := append(pair, array[i], array[j])
					quadruplets = append(quadruplets, newQaudruplet)
				}
			}
		}
		for k := 0; k < i; k++ {
			sum := array[i] + array[k]
			allPairs[sum] = append(allPairs[sum], []int{array[i], array[k]})
		}
	}
	return quadruplets
}
