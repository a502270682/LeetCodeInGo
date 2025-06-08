package sort

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func heapify(nums []int, i, arrLen int) {
	leftSon, rightSon := 2*i+1, 2*i+2
	largest := i
	if leftSon < arrLen && nums[leftSon] > nums[largest] {
		largest = leftSon
	}
	if rightSon < arrLen && nums[rightSon] > nums[largest] {
		largest = rightSon
	}
	if i != largest {
		swap(nums, i, largest)
		heapify(nums, largest, arrLen)
	}
}

func BuildMaxHeap(nums []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(nums, i, arrLen)
	}
}

func heapSort(nums []int) []int {
	arrLen := len(nums)
	BuildMaxHeap(nums, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		swap(nums, 0, i)
		arrLen -= 1
		heapify(nums, 0, arrLen)
	}
	return nums
}
