package leetcode

func singleNumber(nums []int) int {
	result := 0

	for _, number := range nums {

		result ^= number
	}

	return result
}
