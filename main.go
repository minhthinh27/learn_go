package main

import (
	"fmt"
)

func main() {
	var nums = []int{3, 2, 3}

	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	mapArr := make(map[int]int)
	for _, v := range nums {
		mapArr[v]++
		if mapArr[v] > (len(nums) / 2) {
			return v
		}
	}

	return 0
}
