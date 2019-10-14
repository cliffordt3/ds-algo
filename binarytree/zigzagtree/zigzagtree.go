package zigzagtree

import (
	"fmt"
)

type Tree struct {
	value string

	left  *Tree
	right *Tree
}

func (root *Tree) ZigZagTraversal() {
	s1, s2 := []*Tree{}, []*Tree{}
	s1 = append(s1, root)
	for len(s1) > 0 || len(s2) > 0 {
		for len(s1) > 0 {
			n := len(s1) - 1
			p := s1[n]
			s1 = s1[:n]
			fmt.Printf(p.value)
			if p.left != nil {
				s2 = append(s2, p.left)
			}
			if p.right != nil {
				s2 = append(s2, p.right)
			}
		}
		for len(s2) > 0 {
			n := len(s2) - 1
			p := s2[n]
			s2 = s2[:n]
			fmt.Printf(p.value)
			if p.right != nil {
				s1 = append(s1, p.right)
			}
			if p.left != nil {
				s1 = append(s1, p.left)
			}
		}
	}
}
