package dp

func maxSubArray(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	preMax := nums[0]
	maxSubSum := nums[0]
	for i := 1; i < length; i++ {
		preMax = max(preMax+nums[i], nums[i])
		if preMax > maxSubSum {
			maxSubSum = preMax
		}
	}
	return maxSubSum
}
