package main

import (
	"fmt"
)

func main() {
	var nums = []int{1, 2, 3, 1, 1, 3}

	fmt.Println(numIdenticalPairs(nums))
}

func numIdenticalPairs(nums []int) int {
	result := 0
	mapNum := make(map[int]int)

	for _, v := range nums {
		result += mapNum[v]
		mapNum[v]++
	}

	return result
}
