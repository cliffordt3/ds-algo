package invertbinarytree

type BinaryTree struct {
	value int

	left  *BinaryTree
	right *BinaryTree
}

func (tree *BinaryTree) InvertBinaryTree() {
	// Write your code here.
	queue := []*BinaryTree{tree}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == nil {
			continue
		}
		current.right, current.left = current.left, current.right
		queue = append(queue, current.left, current.right)

	}
}
