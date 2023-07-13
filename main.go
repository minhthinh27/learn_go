package main

import (
	"fmt"
)

func main() {
	// var key = "the quick brown fox jumps over the lazy dog"
	// var message = "vkbs bs t suepuv"
	// var nums = []int{0, 1, 4, 6, 7, 10}
	// var diff = 3
	var rings = "B0B6G0R6R0R6G9"
	//var words = []string{"cd", "ac", "dc", "ca", "zz"}
	fmt.Println(countPoints(rings))
}

func countPoints(rings string) int {
	mapRings := make(map[string]map[string]bool)
	result := 0

	for i := 0; i < len(rings); i = i + 2 {
		_, ok := mapRings[string(rings[i+1])]
		if !ok {
			mapRings[string(rings[i+1])] = make(map[string]bool)
			mapRings[string(rings[i+1])][string(rings[i])] = false
		} else {
			_, ok2 := mapRings[string(rings[i+1])][string(rings[i])]
			if !ok2 {
				mapRings[string(rings[i+1])][string(rings[i])] = false
				if len(mapRings[string(rings[i+1])]) == 3 {
					result++
				}
			}
		}
	}
	return result
}
