package arr

func productExceptSelf(nums []int) []int {
	length := len(nums)
	// 设前缀积L和后缀积R
	L := make([]int, length)
	R := make([]int, length)
	l, r := 1, length-2
	L[0], R[length-1] = 1, 1
	for l = 1; l < length; l++ {
		L[l] = nums[l-1] * L[l-1]
		if r >= 0 {
			R[r] = nums[r+1] * R[r+1]
			r--
		}
	}
	ret := make([]int, length)
	for i := 0; i < length; i++ {
		ret[i] = L[i] * R[i]
	}
	return ret
}
