package lca

type BinaryTree struct {
	value       int
	left, right *BinaryTree
}

func LeastCommonAncestor(current, A, B *BinaryTree) *BinaryTree {
	if current == nil {
		return nil
	}
	if current == A || current == B {
		return current
	}
	left := LeastCommonAncestor(current.left, A, B)
	right := LeastCommonAncestor(current, A, B)

	if left != nil && right != nil {
		return current
	}
	if left == nil {
		return right
	}
	return left

}
