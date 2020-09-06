package problem

/*
根据 百度百科 ，生命游戏，简称为生命，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。

给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态：1 即为活细胞（live），或 0 即为死细胞（dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：

如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
根据当前状态，写一个函数来计算面板上所有细胞的下一个（一次更新后的）状态。下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。



示例：

输入：
[
  [0,1,0],
  [0,0,1],
  [1,1,1],
  [0,0,0]
]
输出：
[
  [0,0,0],
  [1,0,1],
  [0,1,1],
  [0,1,0]
]


进阶：

你可以使用原地算法解决本题吗？请注意，面板上所有格子需要同时被更新：你不能先更新某些格子，然后使用它们的更新后的值再更新其他格子。
本题中，我们使用二维数组来表示面板。原则上，面板是无限的，但当活细胞侵占了面板边界时会造成问题。你将如何解决这些问题？

*/

/*
题解
其实就是先统计细胞的状态，然后再统一更新
譬如：
当细胞存活时，使细胞死亡的点在于：
如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
当我们检测到这类细胞时，将这类细胞的状态更新为2
当细胞死亡时，使细胞复活的点在于：
如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
当我们监测到这类细胞时，将这类细胞的状态更新为1

那么相应的，监测细胞周围活细胞的条件为，当身边活细胞状态为>0时，可知它是活细胞

这样确认完所有状态后，下一次一次性更新所有状态即可
*/
func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			temp := countCellsAlive(board, i, j)
			if board[i][j] == 0 {
				if temp == 3 {
					// 复活的记录一个状态
					board[i][j] = -1
				}
			} else {
				if temp < 2 || temp > 3 {
					// 从活变死的记录一个状态
					board[i][j] = 2
				}
			}
		}
	}
	// 统一更新
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			}
			if board[i][j] == -1 {
				board[i][j] = 1
			}
		}
	}
}
func countCellsAlive(board [][]int, i, j int) int {
	res := 0
	for m := -1; m <= 1; m++ {
		for n := -1; n <= 1; n++ {
			if 0 <= (i+m) && (i+m) < len(board) && (j+n) >= 0 && (j+n) < len(board[0]) && board[i+m][j+n] > 0 {
				res++
			}
		}
	}
	if board[i][j] == 1 {
		res--
	}
	return res
}