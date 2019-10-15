package riversizes

func RiverSizes(matrix [][]int) []int {
	sizes := []int{}
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}
	for i := range matrix {
		for j := range matrix[i] {
			if visited[i][j] {
				continue
			}
			sizes = traverseNode(i, j, matrix, visited, sizes)
		}
	}
	return sizes
}

func traverseNode(i, j int, matrix [][]int, visited [][]bool, sizes []int) []int {
	currentRiverSize := 0
	nodesToExplore := [][]int{{i, j}}
	for len(nodesToExplore) > 0 {
		currentNode := nodesToExplore[0]
		nodesToExplore := nodesToExplore[1:]
		i, j := currentNode[0], currentNode[1]
		if visited[i][j] {
			continue
		}
		visited[i][j] = true
		if matrix[i][j] == 0 {
			continue
		}
		currentRiverSize++
		unvisitedNodes := getUnvisitedNodes(i, j, matrix, visited)
		for _, neighbour := range unvisitedNodes {
			nodesToExplore = append(nodesToExplore, neighbour)
		}

	}
	if currentRiverSize > 0 {
		sizes = append(sizes, currentRiverSize)
	}

	return sizes
}

func getUnvisitedNodes(i, j int, matrix [][]int, visited [][]bool) [][]int {
	unvisitedNodes := [][]int{}
	if i > 0 && !visited[i-1][j] {
		unvisitedNodes = append(unvisitedNodes, []int{i - 1, j})
	}
	if i < (len(matrix)-1) && !visited[i+1][j] {
		unvisitedNodes = append(unvisitedNodes, []int{i + 1, j})
	}
	if j > 0 && !visited[i][j-1] {
		unvisitedNodes = append(unvisitedNodes, []int{i, j - 1})
	}
	if j < (len(matrix[0])-1) && !visited[i][j+1] {
		unvisitedNodes = append(unvisitedNodes, []int{i, j + 1})
	}
	return unvisitedNodes
}
