package main

import (
	"fmt"
)

func main() {
	var nums = []int{1, 3}
	var k = 3

	fmt.Println(countKDifference(nums, k))
}

func countKDifference(nums []int, k int) int {
	mapNums := make(map[int]int)
	result := 0

	for _, v := range nums {
		mapNums[v]++
	}

	for _, v := range nums {
		if val, ok := mapNums[v-k]; ok {
			result += val
		}
	}

	return result
}
