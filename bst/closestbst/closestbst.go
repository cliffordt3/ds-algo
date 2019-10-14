package closestbst

import "math"

type BST struct {
	value int

	left  *BST
	right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	current := tree
	diff := math.MaxInt32
	closest := 0
	for current != nil {
		currentDiff := abs(target - current.value)
		if currentDiff < diff {
			diff = currentDiff
			closest = current.value
		}
		if target == current.value {
			return current.value
		} else if target > current.value {
			current = current.right
		} else if target < current.value {
			current = current.left
		}

	}
	return closest
	// Write your code here.
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
