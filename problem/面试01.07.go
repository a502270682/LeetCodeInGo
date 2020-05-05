package problem

/*
给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。

不占用额外内存空间能否做到？

 

示例 1:

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
示例 2:

给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]

 */
// 用翻转代替旋转
func rotate(matrix [][]int)  {
	if len(matrix) == 0 {
		return
	}
	// 纵向翻转
	for y := 0; y < len(matrix)/2; y ++ {
		matrix[y] , matrix[len(matrix)-1-y] = matrix[len(matrix)-1-y], matrix[y]
	}
	// 沿主对角线对折
	// 再按主对角线进行元素翻转
	for y := 0; y < len(matrix); y++ {
		/*
		   这里需要注意，如果用 x<N 这个条件，会导致对角线元素翻转两次，最终恢复原状，所以要 x<y
		   使得翻转的过程只进行一次，即
		   y=0时，不进行翻转
		   y=1时，只进行(1,0)和(0,1)元素翻转
		   y=2时，(2,0)和(0,2)，(2,1)和(1,2)翻转
		*/
		for x := 0; x < y; x++ {
			matrix[y][x], matrix[x][y] = matrix[x][y], matrix[y][x]
		}
	}
}
