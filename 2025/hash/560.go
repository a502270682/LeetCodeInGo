package hash

func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	hash := map[int]int{}
	for start := 0; start < len(nums); start++ {
		hash[pre] += 1
		pre += nums[start]
		if c, ok := hash[pre-k]; ok {
			count += c
		}
	}
	return count
}
