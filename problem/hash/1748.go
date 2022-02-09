package hash

import "sort"

/*
给你一个整数数组nums。数组中唯一元素是那些只出现恰好一次的元素。

请你返回 nums中唯一元素的 和。



示例 1：

输入：nums = [1,2,3,2]
输出：4
解释：唯一元素为 [1,3] ，和为 4 。
示例 2：

输入：nums = [1,1,1,1,1]
输出：0
解释：没有唯一元素，和为 0 。
示例 3 ：

输入：nums = [1,2,3,4,5]
输出：15
解释：唯一元素为 [1,2,3,4,5] ，和为 15 。


提示：

1 <= nums.length <= 100
1 <= nums[i] <= 100

*/

func sumOfUnique(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	sum := 0
	if nums[1] != nums[0] {
		sum += nums[0]
	}
	if nums[len(nums)-1] != nums[len(nums)-2] {
		sum += nums[len(nums)-1]
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			sum += nums[i]
		} else {
			continue
		}
	}
	return sum
}
