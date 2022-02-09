package hash

/*
给你一个整数数组nums和一个整数k，请你返回数对(i, j)的数目，满足i < j且|nums[i] - nums[j]| == k。

|x|的值定义为：

如果x >= 0，那么值为x。
如果x < 0，那么值为-x。


提示：

1 <= nums.length <= 200
1 <= nums[i] <= 100
1 <= k <= 99

*/

func countKDifference(nums []int, k int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]-nums[j] == k || nums[j]-nums[i] == k {
				ret++
			}
		}
	}
	return ret
}

// 空间换时间， hash表
func countKDifference2(nums []int, k int) (ans int) {
	cnt := map[int]int{}
	for _, num := range nums {
		ans += cnt[num-k] + cnt[num+k]
		cnt[num]++
	}
	return
}
