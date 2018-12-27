package Problem0007

import "math"

func reverse(x int) int {
	flag := 1
	if x < 0 {
		flag *= -1
		x = x * -1
	}
	result := 0
	for x > 0 {
		temp := x % 10
		result = result*10 + temp
		x = x / 10
	}
	result = result * flag
	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}
	return result
}
