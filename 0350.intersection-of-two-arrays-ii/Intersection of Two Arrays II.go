package _350_intersection_of_two_arrays_ii

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 =  nums2, nums1
	}

	uniqueMap := map[int]int{}
	for _, i := range nums1 {
		if _, ok := uniqueMap[i]; ok {
			uniqueMap[i] += 1
		} else {
			uniqueMap[i] = 1
		}
	}

	var response []int
	for _, j := range nums2 {
		if num, ok := uniqueMap[j]; ok && num > 0 {
			uniqueMap [j] = num - 1
			response = append(response, j)
		}
	}
	return response
}

