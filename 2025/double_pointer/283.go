package double_pointer

func moveZeroes(nums []int) {
	zeroP, numP := 0, 0
	for zeroP <= numP && numP < len(nums) {
		if nums[numP] == 0 {
			numP++
		} else {
			nums[numP], nums[zeroP] = nums[zeroP], nums[numP]
			zeroP++
			numP++
		}
	}
}
