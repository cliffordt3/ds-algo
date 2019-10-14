package iterativeinorder

type BinaryTree struct {
	value  int
	left   *BinaryTree
	right  *BinaryTree
	parent *BinaryTree
}

func (tree *BinaryTree) IterativeInOrderTraversal(callback func(int)) {

	var previous, next *BinaryTree
	current := tree
	for current != nil {
		if previous == nil || previous == current.parent {
			if current.left != nil {
				next = current.left
			} else {
				callback(current.value)
				if current.right != nil {
					next = current.right
				} else {
					next = current.parent
				}
			}
		} else if previous == current.left {
			callback(current.value)
			if current.right != nil {
				next = current.right
			} else {
				next = current.parent
			}
		} else {
			next = current.parent
		}
		previous = current
		current = next

	}
}
