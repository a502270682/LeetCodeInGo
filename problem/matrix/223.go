package matrix

/*
给你 二维 平面上两个 由直线构成的 矩形，请你计算并返回两个矩形覆盖的总面积。

每个矩形由其 左下 顶点和 右上 顶点坐标表示：

第一个矩形由其左下顶点 (ax1, ay1) 和右上顶点 (ax2, ay2) 定义。
第二个矩形由其左下顶点 (bx1, by1) 和右上顶点 (bx2, by2) 定义。

 */

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	leftX := max(ax1, bx1)
	rightX := min(ax2, bx2)
	upY := min(ay2, by2)
	downY := max(ay1, by1)
	areaA, areaB := (ay2-ay1) * (ax2-ax1), (by2-by1) * (bx2-bx1)
	if leftX >= rightX || downY >= upY {
		return areaB + areaA
	}
	return areaB + areaA - (rightX-leftX)*(upY-downY)
}

func max(x , y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x , y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}