package main

import (
	"fmt"
)

func main() {
	var nums = 746627324245

	fmt.Println(isHappy(nums))
}

func isHappy(n int) bool {
	switch n {
	case 1, 7:
		return true
	case 2, 3, 4, 5, 6, 8, 9:
		return false
	default:
		return isHappy(caculateSum(n))
	}
}

func caculateSum(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}

	return sum
}
