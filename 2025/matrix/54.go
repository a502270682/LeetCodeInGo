package matrix

func spiralOrder(matrix [][]int) []int {
	//colTag, rowTag := 1, 1
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	target := len(matrix) * len(matrix[0])
	ret := make([]int, target)
	count := 0
	for {
		for j := left; j <= right; j++ {
			ret[count] = matrix[up][j]
			count++
		}
		up++
		if breakCondition(count, target) {
			break
		}
		for i := up; i <= down; i++ {
			ret[count] = matrix[i][right]
			count++
		}
		right--
		if breakCondition(count, target) {
			break
		}
		for j := right; j >= left; j-- {
			ret[count] = matrix[down][j]
			count++
		}
		down--
		if breakCondition(count, target) {
			break
		}
		for i := down; i >= up; i-- {
			ret[count] = matrix[i][left]
			count++
		}
		left++
		if breakCondition(count, target) {
			break
		}
	}
	return ret
}

func breakCondition(count, target int) bool {
	return count == target
}
