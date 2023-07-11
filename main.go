package main

import (
	"fmt"
)

func main() {
	// var allowed = "abc"
	// var words = []string{"a", "b", "c", "ab", "ac", "bc", "abc"}
	var nums = []int{0, 1, 4, 6, 7, 10}
	var diff = 3
	fmt.Println(arithmeticTriplets(nums, diff))
}

func arithmeticTriplets(nums []int, diff int) int {
	result := 0
	mapNums := make(map[int]int, len(nums))

	for i, v := range nums {
		mapNums[v] = i
	}

	i, j := 0, 1

	for j < len(nums) {
		if nums[j]-nums[i] == diff {
			if _, ok := mapNums[diff+nums[j]]; ok {
				result++
			}
			i++
		}

		if nums[j]-nums[i] > diff {
			i++
		}

		if nums[j]-nums[i] < diff {
			j++
		}
	}

	return result
}
