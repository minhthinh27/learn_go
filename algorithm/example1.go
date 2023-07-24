package algorithm

func FindDifference(nums1 []int, nums2 []int) [][]int {
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

func AreOccurrencesEqual(s string) bool {
	mapS := make(map[rune]int)

	for _, v := range s {
		mapS[v]++
	}

	for _, v := range mapS {
		if v != len(s)/len(mapS) {
			return false
		}
	}

	return true
}

func SumOfUnique(nums []int) int {
	result := 0
	mapNums := make(map[int]int)

	for _, v := range nums {
		mapNums[v]++
	}

	for k, v := range mapNums {
		if v == 1 {
			result += k
		}
	}

	return result
}
