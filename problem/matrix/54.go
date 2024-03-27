package matrix

/*
给你一个 m 行 n 列的矩阵matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

*/

func spiralOrder(matrix [][]int) []int {
	ret := make([]int, 0)
	left, right, up, down := 0, len(matrix[0])-1, 0, len(matrix)-1
	for left <= right && up <= down {
		for j := left; j <= right; j++ {
			ret = append(ret, matrix[up][j])
		}
		up++
		if up > down {
			break
		}
		for i := up; i <= down; i++ {
			ret = append(ret, matrix[i][right])
		}
		right--
		if left > right {
			break
		}
		for j := right; j >= left; j-- {
			ret = append(ret, matrix[down][j])
		}
		down--
		if up > down {
			break
		}
		for i := down; i >= up; i-- {
			ret = append(ret, matrix[i][left])
		}
		left++
	}
	return ret
}
