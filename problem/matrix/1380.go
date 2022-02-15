package matrix

/*
给你一个 m * n 的矩阵，矩阵中的数字 各不相同 。请你按 任意 顺序返回矩阵中的所有幸运数。

幸运数是指矩阵中满足同时下列两个条件的元素：

在同一行的所有元素中最小
在同一列的所有元素中最大

示例 1：

输入：matrix = [[3,7,8],[9,11,13],[15,16,17]]
输出：[15]
解释：15 是唯一的幸运数，因为它是其所在行中的最小值，也是所在列中的最大值。
示例 2：

输入：matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
输出：[12]
解释：12 是唯一的幸运数，因为它是其所在行中的最小值，也是所在列中的最大值。
示例 3：

输入：matrix = [[7,8],[1,2]]
输出：[7]
*/

func luckyNumbers(matrix [][]int) []int {
	row := len(matrix)
	if row < 1 {
		return nil
	}
	col := len(matrix[0])
	var ret []int

	for i := 0; i < row; i++ {
		thisJ := 0
		rowMin := matrix[i][0]
		for j := thisJ; j < col; j++ {
			if matrix[i][j] < rowMin {
				rowMin = matrix[i][j]
				thisJ = j
			}
		}
		colMax := rowMin
		for thisI := 0; thisI < row; thisI++ {
			if matrix[thisI][thisJ] > colMax {
				colMax = matrix[thisI][thisJ]
			}
		}
		if colMax == rowMin {
			ret = append(ret, rowMin)
		}
	}
	return ret
}
