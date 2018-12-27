package Problem136

func singleNumber(nums []int) int {
	disNum := false
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j {
				if nums[i] == nums[j] {
					disNum = true
				}
			}
		}
		if !disNum {
			return nums[i]
		}
		disNum = false
	}
	return 0
}
