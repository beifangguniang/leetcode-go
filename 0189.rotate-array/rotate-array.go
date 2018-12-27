package Problem0189

func rotate(nums []int, k int) {
	if len(nums) == 1 {
		return
	}
	n := len(nums)
	if k > n {
		k %= n
	}
	if k == 0 || k == n {
		return
	}

	rotating(nums, 0, n-1)
	rotating(nums, k, n-1)
	rotating(nums, 0, k-1)
}

func rotating(nums []int, i int, j int) {
	for i < j {
		nums[j], nums[i] = nums[i], nums[j]
		i++
		j--
	}
}

// 三次翻转
// nums 整体翻转
// nums[:k] 翻转
// nums[k:] 翻转

// 输入: [1,2,3,4,5,6,7] 和 k = 3
// 输出: [5,6,7,1,2,3,4]

// [7,6,5,4,3,2,1]
// [7,6,5,1,2,3,4]
// [5,6,7,1,2,3,4]
