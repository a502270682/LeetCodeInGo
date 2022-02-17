package erfen

/*
给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。

请你找出并返回只出现一次的那个数。

你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。



示例 1:

输入: nums = [1,1,2,3,3,4,4,8,8]
输出: 2
示例 2:

输入: nums =  [3,3,7,7,10,11,11]
输出: 10


提示:

1 <= nums.length <= 105
0 <= nums[i]<= 105

*/
func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return 0
	}

	left, right := 0, len(nums)-1
	for right >= left {
		mid := (right-left)/2 + left
		if mid > 0 && mid < len(nums)-1 && (nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1]) {
			return nums[mid]
		} else if mid == len(nums)-1 && nums[mid] != nums[mid-1] {
			return nums[mid]
		} else if mid == 0 && nums[mid] != nums[mid+1] {
			return nums[mid]
		} else if mid > 0 && nums[mid-1] == nums[mid] {
			if mid%2 != 0 {
				left = mid + 1
			} else {
				right = mid
			}
			continue
		} else if mid < len(nums)-1 && nums[mid+1] == nums[mid] {
			if mid%2 != 0 {
				right = mid
			} else {
				left = mid + 1
			}
			continue
		} else {
			break
		}
	}
	return 0
}
