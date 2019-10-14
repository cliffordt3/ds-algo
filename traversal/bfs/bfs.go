package bfs

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) BreadthFirstSearch(array []string) []string {
	queue := []*Node{n}
	for len(queue) > 0 {
		n := queue[0]
		array = append(array, n.Name)
		for _, child := range n.Children {
			queue = append(queue, child)
		}
		queue = queue[1:]
	}
	return array
}
