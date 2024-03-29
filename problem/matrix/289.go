package matrix

func gameOfLife(board [][]int) {
	left, right, up, down := 0, len(board[0])-1, 0, len(board)-1
	tmp := make([][]int, down+1)
	for i := 0; i <= down; i++ {
		row := make([]int, right+1)
		for j := 0; j <= right; j++ {
			row[j] = board[i][j]
		}
		tmp[i] = row
	}
	countLiveNeighbor := func(i, j int) int {
		count := 0
		if i > up {
			if tmp[i-1][j] == 1 {
				count++
			}
			if j > left && tmp[i-1][j-1] == 1 {
				count++
			}
			if j < right && tmp[i-1][j+1] == 1 {
				count++
			}
		}
		if j > left && tmp[i][j-1] == 1 {
			count++
		}
		if j < right && tmp[i][j+1] == 1 {
			count++
		}
		if i < down {
			if tmp[i+1][j] == 1 {
				count++
			}
			if j > left && tmp[i+1][j-1] == 1 {
				count++
			}
			if j < right && tmp[i+1][j+1] == 1 {
				count++
			}
		}
		return count
	}
	for i := 0; i <= down; i++ {
		for j := 0; j <= right; j++ {
			this := board[i][j]
			if count := countLiveNeighbor(i, j); count < 2 {
				this = 0
			} else if count == 2 {
				// do nothing
			} else if count == 3 {
				this = 1
			} else if count > 3 {
				this = 0
			}
			board[i][j] = this
		}
	}
	return
}
