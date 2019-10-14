package haspathsum

type BinaryTree struct {
	value       int
	left, right *BinaryTree
}

func HasPathSum(tree *BinaryTree, target int) bool {
	return helper(tree, target, 0)
}

func helper(tree *BinaryTree, target, partialSum int) bool {
	if tree == nil {
		return false
	}
	partialSum += tree.value
	if tree.left == nil && tree.right == nil {
		return partialSum == target
	}
	return helper(tree.left, target, partialSum) || helper(tree.right, target, partialSum)

}
