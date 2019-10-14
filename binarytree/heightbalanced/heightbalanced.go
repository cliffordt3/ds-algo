package heightbalanced

type BinaryTree struct {
	value       int
	left, right *BinaryTree
}
type BinaryTreeStatusWithHeight struct {
	balanced bool
	height   int
}

func NewBinaryStatusWithHeight(balanced bool, height int) *BinaryTreeStatusWithHeight {
	return &BinaryTreeStatusWithHeight{
		balanced: balanced,
		height:   height,
	}
}

func HeightBalanced(tree *BinaryTree) bool {
	balancedStatus := helper(tree)
	return balancedStatus.balanced
}

func helper(tree *BinaryTree) *BinaryTreeStatusWithHeight {
	if tree == nil {
		return NewBinaryStatusWithHeight(true, -1)
	}
	leftStatus := helper(tree.left)
	if !leftStatus.balanced {
		return leftStatus
	}
	rightStatus := helper(tree.right)
	if !rightStatus.balanced {
		return rightStatus
	}
	height := max(rightStatus.height, leftStatus.height) + 1
	if abs(leftStatus.height-rightStatus.height) > 1 {
		return NewBinaryStatusWithHeight(false, height)
	}
	return NewBinaryStatusWithHeight(true, height)

}
func max(n, m int) int {
	if n > m {
		return n
	}
	return m

}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
