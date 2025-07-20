package backtrack

func exist(board [][]byte, word string) bool {
	if word == "" || len(board) == 0 {
		return false
	}
	rowLength := len(board) - 1
	colLength := len(board[0]) - 1
	first := word[0]
	var dfs func(index, x, y int) bool
	dfs = func(index, x, y int) bool {
		if x < 0 || x > rowLength {
			return false
		}
		if y < 0 || y > colLength {
			return false
		}
		if len(word) == index {
			return true
		}
		if x > 0 {
			// 左
			if board[x-1][y] == word[index] {
				dfs(index+1, x-1, y)
			}
		}
		if x < rowLength {
			// 右
			if board[x+1][y] == word[index] {
				dfs(index+1, x+1, y)
			}
		}
		if y > 0 {
			// 上
			if board[x][y-1] == word[index] {
				dfs(index+1, x, y-1)
			}
		}
		if y < colLength {
			// 下
			if board[x][y+1] == word[index] {
				dfs(index+1, x, y+1)
			}
		}
		return false
	}
	for row, colArr := range board {
		for col, value := range colArr {
			if value == first {
				if dfs(1, row, col) {
					return true
				}
			}
		}
	}
	return false
}
