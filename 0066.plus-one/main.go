package Problem0066

// [9] 			     [1,0]
// [1,2,3]		     [1,2,4]
// [9,9,9,9]		[1,0,0,0,0]
//[1,1,9,9]			[1,2,0,0]
func plusOnes(digits []int) []int {
	digitsPlus := make([]int, len(digits)+1)
	if digits[0] > 8 {
		for i := len(digits) - 1; i > 0; i-- {
			if digits[i] > 8 {
				digitsPlus[i] = 0
			} else {
				digits[i] += 1
				return digits
			}
		}
		digitsPlus[0] = 1
		return digitsPlus
	}

	if digits[len(digits)-1] <= 8 {
		digits[len(digits)-1] += 1
		return digits
	} else {
		count := 1
		for i := len(digits) - 1; i >= 0; i-- {
			if digits[i] <= 8 {
				digits[i] += 1
				return digits
			} else {
				digits[i] = 0
				count++
			}
		}
	}
	return digits
}
