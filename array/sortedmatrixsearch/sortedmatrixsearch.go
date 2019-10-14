package sortedmatrixsearch

func SearchInSortedMatrix(matrix [][]int, target int) []int {
	row := 0 
	col := len(matrix[0]) - 1
	for row < len(matrix) && col >=0 {
		if matrix[row][col] > target {
			col--
		}else if matrix[row][col] < target {
			row++
		}else{
			return []int{row,col}
		}
	}
	return []int{-1,-1}
	// Write your code here.
}
