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
	var names = []string{"Alice", "Bob", "Marry"}
	var heights = []int{155, 185, 160}
	fmt.Println(sortPeople(names, heights))
}

func sortPeople(names []string, heights []int) []string {

	mapHeight := make(map[int]string)

	for i, v := range heights {
		mapHeight[v] = names[i]
	}

	start, end := 0, len(heights)-1

	for start < end {
		i := start

		for i < end {
			if heights[i+1] > heights[start] {
				heights[start], heights[i+1] = heights[i+1], heights[start]
			}
			i++
		}

		start++
	}

	result := []string{}
	for _, val := range heights {
		result = append(result, mapHeight[val])
	}

	return result
}
