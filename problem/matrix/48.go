package matrix

// x(i, j) -> x(j, n-1-i)

func rotate(matrix [][]int) {
	n := len(matrix)
	round := n / 2
	movNums := n - 1
	for i := 0; i < round; i++ {
		for j := 0; j < (n+1)/2; j++ {
			tmp := matrix[i][j]
			// 左下 -> 左上
			matrix[i][j] = matrix[movNums-j][i]
			// 右下 -> 左下
			matrix[movNums-j][i] = matrix[movNums-i][movNums-j]
			// 右上 -> 右下
			matrix[movNums-i][movNums-j] = matrix[j][movNums-i]
			// 左下 -> 右上
			matrix[j][movNums-i] = tmp
		}
	}
}
