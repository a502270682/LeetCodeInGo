package double_pointer

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		cur := (right - left) * min(height[left], height[right])
		if cur > maxArea {
			maxArea = cur
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
