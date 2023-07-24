package main

import (
	"fmt"
	"learn_go/algorithm"
)

func main() {
	var nums1 = []int{5, 1, 5, 2, 5, 3, 5, 4}
	// var nums2 = []int{2, 4, 6}
	//var s = "vvvvvvvvvvvvvvvvvvv"
	result := algorithm.RepeatedNTimes(nums1)
	fmt.Println(result)
}
