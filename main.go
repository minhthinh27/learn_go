package main

import (
	"fmt"
	"learn_go/algorithm"
)

func main() {
	var nums1 = []int{1, 2, 3, 2}
	// var nums2 = []int{2, 4, 6}
	//var s = "vvvvvvvvvvvvvvvvvvv"
	result := algorithm.SumOfUnique(nums1)
	fmt.Println(result)
}
