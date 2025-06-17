package double_pointer

import "sort"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1

		for left < right {
			if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				for left < len(nums) && left > i+1 && nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}
	return res
}
