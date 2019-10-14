package binarysearch

func BinarySearch(array []int, target int) int {
	return binarySearchHelper(array, target, 0, len(array)-1)
}

func binarySearchHelper(array []int, target int, start int, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	potentialMatch := array[mid]
	if target == potentialMatch {
		return mid
	} else if target < potentialMatch {
		return binarySearchHelper(array, target, start, mid-1)
	}
	return binarySearchHelper(array, target, mid+1, end)
}
