package main

import "fmt"

func main() {
	test := []int{2, 2, 2, 1, 1, 2}
	sortColors(test)
	// fmt.Println(sortColors(test))
}

// func sortColors(nums []int) {
// 	mapColor := make(map[int]int, len(nums))

// 	for _, v := range nums {
// 		mapColor[v]++
// 	}

// 	fmt.Println(mapColor)

// 	// red: 0 , white: 1 , blue: 2
// 	nameColor := 0
// 	i := 0
// 	for i < len(nums) {
// 		value, ok := mapColor[nameColor]
// 		if !ok || value == 0 {
// 			nameColor++
// 		} else if value > 0 {
// 			nums[i] = nameColor
// 			mapColor[nameColor] = value - 1
// 			i++
// 		}
// 	}

// 	// for i, _ := range nums {
// 	// 	if value, ok := mapColor[nameColor]; ok && value > 0 {
// 	// 		mapColor[nameColor]--
// 	// 		nums[i] = nameColor
// 	// 	} else {
// 	// 		nameColor += 1
// 	// 	}
// 	// }

// 	fmt.Println(nums)
// }

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	var start, end = 0, len(nums) - 1
	for i := 0; i <= end; i++ {
		// swap current to end
		if nums[i] == 2 {
			nums[end], nums[i] = nums[i], nums[end]
			end--
			i--
		} else if nums[i] == 0 {
			nums[start], nums[i] = nums[i], nums[start]
			start++
		}
	}

	fmt.Println(nums)
}
