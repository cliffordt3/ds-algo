package apartmenthunting

import "math"

type Block map[string]bool

func ApartmentHunting(blocks []Block, reqs []string) int {
	// Write your code here.
	maxDistanceAtBlocks := make([]int, len(blocks))
	for i := range blocks {
		maxDistanceAtBlocks[i] = -1
		for _, req := range reqs {
			closestReqDistance := math.MaxInt32
			for j := range blocks {
				if blocks[j][req] {
					closestReqDistance = min(closestReqDistance, distanceBetween(i, j))
				}
			}
			maxDistanceAtBlocks[i] = max(maxDistanceAtBlocks[i], closestReqDistance)
		}
	}
	return getBestBlockIdx(maxDistanceAtBlocks)
}

func getBestBlockIdx(array []int) int {
	bestIdx := 0
	bestBlock := math.MaxInt32
	for i := range array {
		if array[i] < bestBlock {
			bestBlock = array[i]
			bestIdx = i
		}
	}
	return bestIdx
}

func distanceBetween(i, j int) int {
	return abs(i - j)
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func min(n, m int) int {
	if n < m {
		return n
	}
	return m
}
func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}

func ApartmentHuntingV2(blocks []Block, reqs []string) int {
	minDistancesForBlocks := [][]int{}
	for _, req := range reqs {
		minDistancesForBlocks = append(minDistancesForBlocks, getMinDistances(blocks, req))
	}
	maxDistanceFromBlocks := getMaxDistanceAtBlocks(blocks, minDistancesForBlocks)
	return getBestBlockIdx(maxDistanceFromBlocks)
}

func getMinDistances(blocks []Block, req string) []int {
	minDistances := make([]int, len(blocks))
	closestReq := math.MaxInt32
	for i := range blocks {
		if blocks[i][req] {
			closestReq = i
		}
		minDistances[i] = distanceBetween(i, closestReq)
	}
	for i := len(blocks) - 1; i > 0; i-- {
		if blocks[i][req] {
			closestReq = i
		}
		minDistances[i] = min(minDistances[i], distanceBetween(i, closestReq))
	}
	return minDistances
}

func getMaxDistanceAtBlocks(blocks []Block, minDistancesFromBlocks [][]int) []int {
	maxDistancesForBlocks := make([]int, len(blocks))
	for i := range blocks {
		minDistancesAtBlock := []int{}
		for _, distances := range minDistancesFromBlocks {
			minDistancesAtBlock = append(minDistancesAtBlock, distances[i])
		}
		maxDistancesForBlocks[i] = maxInArr(minDistancesAtBlock)
	}
	return maxDistancesForBlocks
}

func maxInArr(array []int) int {
	max := array[0]
	for _, val := range array {
		if val > max {
			max = val
		}
	}
	return max
}
