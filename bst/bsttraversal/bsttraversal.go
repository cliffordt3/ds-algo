package bsttraversal

type BST struct {
	value int

	left  *BST
	right *BST
}

func (tree *BST) InOrderTraverse(array []int) []int {
	if tree.left != nil {
		array = tree.left.InOrderTraverse(array)
	}
	array = append(array, tree.value)
	if tree.right != nil {
		array = tree.right.InOrderTraverse(array)
	}
	return array
	// Write your code here.
}

func (tree *BST) PreOrderTraverse(array []int) []int {

	array = append(array, tree.value)
	if tree.left != nil {
		array = tree.left.PreOrderTraverse(array)
	}
	if tree.right != nil {
		array = tree.right.PreOrderTraverse(array)
	}

	return array
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	if tree.left != nil {
		array = tree.left.PostOrderTraverse(array)
	}
	if tree.right != nil {
		array = tree.right.PostOrderTraverse(array)
	}
	array = append(array, tree.value)
	return array
}
