package stack

func decodeString(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, s[i])
		} else {
			// 遇到 '[' 时进行处理

			// 先取出[]这里的内容
			var tmp []byte
			j := len(stack) - 1
			for j >= 0 && stack[j] != '[' {
				tmp = append(tmp, stack[j])
				j--
			}
			// 弹出刚才收集到的内容，包括'['
			stack = stack[:j]

			// 处理数字
			num := 0
			base := 1
			m := j - 1
			for m >= 0 && stack[m] >= '0' && stack[m] <= '9' {
				num += int(stack[m]-'0') * base
				base *= 10
				m--
			}
			stack = stack[:m+1]

			// 翻转tmp，因为取出来是反的
			for j = 0; j < len(tmp)/2; j++ {
				tmp[j], tmp[len(tmp)-1-j] = tmp[len(tmp)-1-j], tmp[j]
			}

			for j = 1; j <= num; j++ {
				stack = append(stack, tmp...)
			}
		}
	}
	return string(stack)
}
