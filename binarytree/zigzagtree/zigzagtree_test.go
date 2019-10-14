package zigzagtree

import (
	"testing"
)

func NewTree(value string) *Tree {
	return &Tree{
		value: value,
	}
}

func (tree *Tree) AddLeftTree(value string) *Tree {
	lchild := &Tree{value: value}
	tree.left = lchild
	return tree
}

func (tree *Tree) AddRightChild(value string) *Tree {
	rchild := &Tree{value: value}
	tree.right = rchild
	return tree
}

var test1 = NewTree("A").AddLeftTree("B").AddLeftTree("D")
test1 = test1.AddRightChild("E")
test1 = test1.AddRightChild("C")

func TestCase1(t *testing.T) {
	test1.ZigZagTraversal()
}

