package main

func maximumNumberOfStringPairs(words []string) int {
	wordM := make(map[string]int)
	couple := 0
	for _, w := range words {
		if count, ok := wordM[w]; ok {
			couple++
			wordM[w] = count - 1
		} else {
			tmp := ""
			for i := len(w) - 1; i >= 0; i-- {
				tmp += string(w[i])
			}
			wordM[tmp] = 1
		}
	}
	return couple
}
