package Problem0217

func containsDuplicate(nums []int) bool {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j {
				if nums[i] == nums[j] {
					return true
				}
			}
		}
	}
	return false
}
