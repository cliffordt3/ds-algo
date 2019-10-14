package zigzagtraversal

func ZigzagTraverse(array [][]int) []int {
	height := len(array) - 1
	width := len(array[0]) - 1
	row, col := 0, 0
	result := []int{}
	goingDown := true

	for !isOutOfBounds(row, col, width, height) {
		result = append(result, array[row][col])
		if goingDown {
			if row == height || col == 0 {
				goingDown = false
				if row == height {
					col++
				} else {
					row++
				}
			} else {
				row++
				col--
			}
		} else {
			if row == 0 || col == width {
				goingDown = true
				if col == width {
					row++
				} else {
					col++
				}
			} else {
				row--
				col++
			}
		}
	}
	return result
}

func isOutOfBounds(row, col, width, height int) bool {
	return (row < 0 || row > height || col < 0 || col > width)
}
