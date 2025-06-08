package _025

func clearStars(s string) string {
	letterStack := make([][]int, 26)
	for i := range letterStack {
		letterStack[i] = make([]int, 0)
	}
	arr := []rune(s)
	for i, c := range arr {
		if c != '*' {
			letterStack[c-'a'] = append(letterStack[c-'a'], i)
		} else {
			for j := range letterStack {
				if len(letterStack[j]) > 0 {
					arr[letterStack[j][len(letterStack[j])-1]] = '*'
					letterStack[j] = letterStack[j][:len(letterStack[j])-1]
					break
				}
			}
		}
	}
	res := ""
	for _, v := range arr {
		if v != '*' {
			res += string(v)
		}
	}
	return res
}
