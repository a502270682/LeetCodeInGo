package stack

func dailyTemperatures(temperatures []int) []int {
	// 从栈底到顶单调递减栈
	indexStack := make([]int, 0)
	res := make([]int, len(temperatures))
	// 从右到左构建单调递减栈
	for i := len(temperatures) - 1; i >= 0; i-- {
		value := temperatures[i]
		if len(indexStack) == 0 {
			indexStack = append(indexStack, i)
			res[i] = 0
		} else {
			if temperatures[indexStack[len(indexStack)-1]] > value {
				// 符合，直接入栈
				// 比他更大的就是下一个数
				res[i] = 1
			} else {
				for len(indexStack) > 0 && temperatures[indexStack[len(indexStack)-1]] <= value {
					// 栈顶元素小于当前元素，需要indexStack.pop()
					indexStack = indexStack[:len(indexStack)-1]
				}
				if len(indexStack) > 0 {
					res[i] = indexStack[len(indexStack)-1] - i
				} else {
					res[i] = 0
				}
			}
			indexStack = append(indexStack, i)
		}
	}
	return res
}
