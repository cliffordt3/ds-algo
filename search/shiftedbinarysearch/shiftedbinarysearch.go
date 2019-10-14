package shiftedbinarysearch

func ShiftedBinarySearch(array []int, target int) int {
	// Write your code here.
	left, right := 0, len(array)-1

	for left <= right {
		mid := (left + right) / 2
		if target == array[mid] {
			return mid
		}
		if array[left] <= array[mid] {
			if target < array[mid] && target >= array[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > array[mid] && target <= array[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
