package stack

func MaxNumber(nums1 []int, nums2 []int, k int) []int {
	return maxNumber(nums1, nums2, k)
}

// 分为两个问题，先各自对nums做获取， 再合并两个结果
// 相当于316 + merge
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	max := []int{}
	for i := 0; i <= k; i++ {
		if i <= len(nums1) && k-i <= len(nums2) {
			this := merge(pickMaxNums(nums1, i), pickMaxNums(nums2, k-i))
			if lexicographicalLess(max, this) {
				max = this
			}
		}
	}
	return max
}

func lexicographicalLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func merge(a, b []int) []int {
	merged := make([]int, len(a)+len(b))
	for i := range merged {
		if lexicographicalLess(a, b) {
			merged[i] = b[0]
			b = b[1:]
		} else {
			merged[i] = a[0]
			a = a[1:]
		}
	}
	return merged
}

func pickMaxNums(arr []int, k int) []int {
	stack := []int{}
	move := len(arr) - k
	for _, i := range arr {
		for len(stack) > 0 && move > 0 && stack[len(stack)-1] < i {
			stack = stack[:len(stack)-1]
			move--
		}
		stack = append(stack, i)
	}
	if move > 0 {
		if move >= len(stack) {
			return []int{}
		} else {
			return stack[:len(stack)-move]
		}
	} else {
		return stack
	}
}
