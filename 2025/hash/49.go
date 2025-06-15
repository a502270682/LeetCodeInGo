package hash

func groupAnagrams(strs []string) [][]string {
	hash := make(map[[26]int][]string)
	for _, str := range strs {
		tmp := [26]int{}
		for _, ch := range str {
			tmp[ch-'a']++
		}
		if res, ok := hash[tmp]; ok {
			res = append(res, str)
			hash[tmp] = res
		} else {
			hash[tmp] = []string{str}
		}
	}
	res := make([][]string, 0)
	for _, v := range hash {
		if len(v) > 0 {
			res = append(res, v)
		}
	}
	return res
}
