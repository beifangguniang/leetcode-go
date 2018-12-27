package main

func moveZeroes(nums []int) {
	i, j := 0, 0
	l := len(nums)
	for j < l {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	for i < l {
		nums[i] = 0
		i++
	}
}
