func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	temp := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if temp == nums[i] {
			continue
		} else {
			nums[count] = nums[i]
			temp = nums[count]
		}
		count++
	}
	return count
}
