package validatebst

import "math"

type BST struct {
	value int

	left  *BST
	right *BST
}

func (tree *BST) Validate() bool {
	return tree.validateBSTHelper(math.MinInt32, math.MaxInt32)
	// Write your code here.
}

func (tree *BST) validateBSTHelper(minValue, maxValue int) bool {
	if tree == nil {
		return true
	}
	if tree.value < minValue || tree.value >= maxValue {
		return false
	}
	leftTreeValid := tree.left.validateBSTHelper(minValue, tree.value)
	rightTreeValid := tree.right.validateBSTHelper(tree.value, maxValue)
	return leftTreeValid && rightTreeValid
}
