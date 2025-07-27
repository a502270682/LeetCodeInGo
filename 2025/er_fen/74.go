package er_fen

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return false
	}
	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)/2
		if matrix[mid/n][mid%n] == target {
			return true
		} else if matrix[mid/n][mid%n] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	col := len(matrix[0])
	nums := make([]int, len(matrix)*col)
	for i, row := range matrix {
		for j, num := range row {
			nums[i*col+j] = num
		}
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
