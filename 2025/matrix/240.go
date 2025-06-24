package matrix

func searchMatrix(matrix [][]int, target int) bool {
	y := len(matrix[0]) - 1
	x := 0
	for y >= 0 && x < len(matrix) {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}
