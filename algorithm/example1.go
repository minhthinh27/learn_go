package algorithm

func findDifference(nums1 []int, nums2 []int) [][]int {
	var result [][]int
	mapNum1 := make(map[int]bool)
	mapNum2 := make(map[int]bool)

	for _, v := range nums1 {
		mapNum1[v] = true
	}

	for _, v := range nums2 {
		mapNum2[v] = true
	}

	result1 := []int{}
	for k := range mapNum1 {
		if _, ok := mapNum2[k]; !ok {
			result1 = append(result1, k)
		}
	}

	result2 := []int{}
	for k := range mapNum2 {
		if _, ok := mapNum1[k]; !ok {
			result2 = append(result2, k)
		}
	}

	result = append(result, result1, result2)

	return result
}
