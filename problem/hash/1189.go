package hash

/*
给你一个字符串text，你需要使用 text 中的字母来拼凑尽可能多的单词"balloon"（气球）。

字符串text 中的每个字母最多只能被使用一次。请你返回最多可以拼凑出多少个单词"balloon"。

*/

func maxNumberOfBalloons(text string) int {
	letter := make([]int, 26)
	for _, this := range text {
		letter[this-'a'] += 1
	}
	minSingle := 10000
	minDouble := 10000
	for _, this := range "balloon" {
		tmp := letter[this-'a']
		if this == 'a' || this == 'b' || this == 'n' {
			if tmp < minSingle {
				minSingle = tmp
			}
		} else if this == 'l' || this == 'o' {
			if tmp < minDouble {
				minDouble = tmp
			}
		}

	}
	if minSingle == 0 {
		return 0
	}
	for i := minSingle; i > 0; i-- {
		if minDouble >= i*2 {
			return i
		}
	}
	return 0
}
