package stack

func removeDuplicateLetters(s string) string {
	bMap := make([]int, 26)
	for _, i := range s {
		bMap[i-'a']++
	}
	stack := []rune{}
	inStack := make([]bool, 26)
	// 从前往后，把栈顶比当前准备入栈的要大的，且出现次数超过一次的去掉
	for _, i := range s {
		if inStack[i-'a'] {
			bMap[i-'a']--
			continue
		}
		for len(stack) > 0 && bMap[stack[len(stack)-1]-'a'] > 1 && stack[len(stack)-1] > i {
			bMap[stack[len(stack)-1]-'a']--
			inStack[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		inStack[i-'a'] = true
	}
	return string(stack)
}
