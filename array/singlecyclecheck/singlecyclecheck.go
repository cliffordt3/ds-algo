package singlecyclecheck

func HasSingleCycle(array []int) bool {
	// Write your code here.
	numVisited := 0
	currentIdx := 0
	for numVisited < len(array) {
		if numVisited > 0 && currentIdx == 0 {
			return false
		}
		numVisited++
		currentIdx = getNextIdx(currentIdx, array)
	}
	return currentIdx == 0
}

func getNextIdx(currentIdx int, array []int) int {
	jump := array[currentIdx]
	nextIdx := (currentIdx + jump) % len(array)
	if nextIdx >= 0 {
		return nextIdx
	}
	return nextIdx + len(array)

}
