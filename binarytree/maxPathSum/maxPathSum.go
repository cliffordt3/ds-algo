package maxPathSum

type BinaryTree struct {
	value       int
	left, right *BinaryTree
}

func MaxPathSum(tree *BinaryTree) int {
	_, maxSum := findMaxsum(tree)
	return maxSum
}

func findMaxsum(tree *BinaryTree) (int, int) {
	if tree == nil {
		return 0, 0
	}
	leftMaxSumAsBranch, leftMaxPathSum := findMaxsum(tree.left)
	rightMaxSumAsBranch, rightMaxPathSum := findMaxsum(tree.right)

	maxChildSumAsBranch := max(leftMaxSumAsBranch, rightMaxSumAsBranch)
	value := tree.value
	maxSumAsBranch := max(maxChildSumAsBranch+value, value)
	maxSumAsRootNode := max(leftMaxSumAsBranch+value+
		rightMaxSumAsBranch, maxSumAsBranch)
	maxPathSum := max(leftMaxPathSum, rightMaxPathSum, maxSumAsRootNode)
	return maxSumAsBranch, maxPathSum
}

func max(n int, m ...int) int {
	for _, val := range m {
		if val > n {
			n = val
		}
	}
	return n
}
