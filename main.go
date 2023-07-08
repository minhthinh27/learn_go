package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nums = []int{0, 1, 2, 4, 5, 7}

	fmt.Println(summaryRanges(nums))
}

func summaryRanges(nums []int) []string {
	result := []string{}
	temp := ""

	for i := 0; i < len(nums); {
		temp += strconv.Itoa(nums[i])

		j := i + 1
		for ; j < len(nums) && nums[j]-nums[j-1] == 1; j++ {
		}

		if j-1 > i {
			temp = temp + "->" + strconv.Itoa(nums[j-1])
		}

		result = append(result, temp)
		temp = ""
		i = j
	}

	return result
}
