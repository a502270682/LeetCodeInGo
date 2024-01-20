package moni

func splitWordsBySeparator(words []string, separator byte) []string {
	ret := make([]string, 0)
	for _, w := range words {
		start := 0
		for i := 0; i < len(w); i++ {
			if w[i] == separator {
				if start < i {
					ret = append(ret, w[start:i])
				}
				start = i + 1
			}
		}
		if start <= len(w)-1 {
			ret = append(ret, w[start:])
		}
	}
	return ret
}
