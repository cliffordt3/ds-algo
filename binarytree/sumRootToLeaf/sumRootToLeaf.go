package sumRootToLeaf

type BinaryTree struct {
	value       int
	left, right *BinaryTree
}

func SumRootToLeaf(tree *BinaryTree) int {
	return helper(tree, 0)
}

func helper(root *BinaryTree, partialSum int) int {
	if root == nil {
		return 0
	}
	partialSum = partialSum*2 + root.value

	if root.left == nil && root.right == nil {
		return partialSum
	}
	return helper(root.left, partialSum) + helper(root.right, partialSum)

}
