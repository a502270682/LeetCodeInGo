package hash

/*
给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 。请你找出所有出现 两次 的整数，并以数组形式返回。

你必须设计并实现一个时间复杂度为 O(n) 且仅使用常量额外空间的算法解决此问题。


示例 1：

输入：nums = [4,3,2,7,8,2,3,1]
输出：[2,3]
示例 2：

输入：nums = [1,1,2]
输出：[1]
示例 3：

输入：nums = [1]
输出：[]
 */
func findDuplicates(nums []int) []int {
	ret := make([]int, 0)
	for idx := range nums {
		// 第一次遍历，把遍历到的数字nums[idx] 和 对应该数字的下标值nums[nums[idx]-1] 比较，
		// 如果不相同，则把nums[idx]移动到该在的下标位置，即交换位置
		// 直到找到与最后一个被放置到对应下标位置相同的数后，将其放在当前idx位，结束for循环
		for nums[idx] != nums[nums[idx]-1] {
			nums[idx], nums[nums[idx]-1] = nums[nums[idx]-1], nums[idx]
		}
	}
	for idx, val := range nums {
		// 第二次遍历，由于对应值已经下标位置一定有nums[idx] = idx+1
		// 所以如果发现当前不相等的，一定是重复出现的数组
		if idx+1 != val {
			ret = append(ret, val)
		}
	}
	return ret
}