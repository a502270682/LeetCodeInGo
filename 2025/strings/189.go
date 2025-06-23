package strings

func rotate(nums []int, k int) {
	if len(nums) <= 1 || k == 0 {
		return
	}
	length := len(nums)
	count := 0
	start := 0
	cur := 0
	tmp := nums[0]
	for count < length {
		cur = (cur + k) % length
		tmp, nums[cur] = nums[cur], tmp
		count++
		if count == length {
			break
		}
		if start == cur {
			start++
			cur = start
			tmp = nums[cur]
		}
	}
}
