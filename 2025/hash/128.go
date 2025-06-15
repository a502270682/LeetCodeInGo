package hash

func longestConsecutive(nums []int) int {
	hash := make(map[int]bool)
	for _, num := range nums {
		hash[num] = true
	}
	maxLength := 0
	for num := range hash {
		if !hash[num-1] {
			curLen := 1
			for hash[num+1] {
				num++
				curLen++
			}
			if curLen > maxLength {
				maxLength = curLen
			}
		}

	}
	return maxLength
}
