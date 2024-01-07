package main

func canConstruct(ransomNote string, magazine string) bool {
	magazineM := make([]int, 26)
	for _, i := range magazine {
		magazineM[i-'a']++
	}
	for _, i := range ransomNote {
		if magazineM[i-'a'] > 0 {
			magazineM[i-'a']--
		} else {
			return false
		}
	}
	return true
}
