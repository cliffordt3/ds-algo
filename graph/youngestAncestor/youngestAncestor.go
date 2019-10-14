package youngestAncestor

type AncestralTree struct {
	Name     string
	Ancestor *AncestralTree
}

func GetYoungestCommonAncestor(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	// Write your code here.
	depth1 := getDepth(top, descendantOne)
	depth2 := getDepth(top, descendantTwo)
	if depth1 > depth2 {
		return backTrackToCommonAncestor(descendantOne, descendantTwo, depth1-depth2)
	}
	return backTrackToCommonAncestor(descendantTwo, descendantOne, depth2-depth1)

}

func getDepth(ancestor, child *AncestralTree) int {
	depth := 0
	for child != ancestor {
		depth++
		child = child.Ancestor
	}
	return depth
}

func backTrackToCommonAncestor(lowerDescendant, higherDescendant *AncestralTree, diff int) *AncestralTree {
	for diff > 0 {
		lowerDescendant = lowerDescendant.Ancestor
		diff--
	}
	for lowerDescendant != higherDescendant {
		lowerDescendant = lowerDescendant.Ancestor
		higherDescendant = higherDescendant.Ancestor
	}
	return lowerDescendant
}
