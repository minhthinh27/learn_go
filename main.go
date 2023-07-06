package main

import "fmt"

func main() {
	test := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//sortColors(test)
	fmt.Println(maxArea(test))
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	length := len(height) - 1
	res := 0

	for i != j {
		temp := 0
		if height[i] < height[j] {
			temp = height[i] * length
			i++
		} else {
			temp = height[j] * length
			j--
		}

		length--

		if temp > res {
			res = temp
		}
	}

	return res
}
