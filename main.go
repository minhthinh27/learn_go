package main

import "fmt"

func main() {
	test := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(test))
}

func sortArrayByParityII(nums []int) []int {
	evenIndex, oddIndex := 0, 1

	result := make([]int, len(nums))

	for _, v := range nums {
		if v%2 == 0 {
			result[evenIndex] = v
			evenIndex += 2
		} else {
			result[oddIndex] = v
			oddIndex += 2
		}
	}

	return result
}

// func sortArrayByParityII(A []int) []int {
// 	even, odd := 0, 1
// 	for even < len(A) && odd < len(A) {
// 		if A[even]%2 != 0 {
// 			A[even], A[odd] = A[odd], A[even]
// 			odd += 2
// 		} else {
// 			even += 2
// 		}
// 	}
// 	return A
// }
