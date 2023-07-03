package main

import (
	"fmt"
)

func main() {
	// var key = "the quick brown fox jumps over the lazy dog"
	// var message = "vkbs bs t suepuv"
	// var nums = []int{0, 1, 4, 6, 7, 10}
	// var diff = 3
	//var rings = "B0B6G0R6R0R6G9"
	// var names = []string{"Alice", "Bob", "Marry"}
	// var heights = []int{155, 185, 160}
	num1 := []int{1, 2, 3}
	num2 := []int{2, 4, 6}
	fmt.Println(findDifference(num1, num2))
}

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

// func findDifference(nums1 []int, nums2 []int) [][]int {
// 	return [][]int{setDifference(nums1, nums2), setDifference(nums2, nums1)}
// }

// func setDifference(a, b []int) []int {
// 	d := [2001]bool{}
// 	ans := []int{}

// 	for _, x := range b {
// 		d[x+1000] = true
// 	}

// 	for _, x := range a {
// 		if !d[x+1000] {
// 			ans = append(ans, x)
// 			d[x+1000] = true // prevent duplicates from being added
// 		}
// 	}

// 	return ans
// }
