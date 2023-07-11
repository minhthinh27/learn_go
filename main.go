package main

import (
	"fmt"
)

func main() {
	// var key = "the quick brown fox jumps over the lazy dog"
	// var message = "vkbs bs t suepuv"
	// var nums = []int{0, 1, 4, 6, 7, 10}
	// var diff = 3
	var words = []string{"cd", "ac", "dc", "ca", "zz"}
	fmt.Println(maximumNumberOfStringPairs(words))
}

func maximumNumberOfStringPairs(words []string) int {
	mapWords := make(map[string]bool)
	result := 0
	for _, str := range words {
		temp := reverseString(str)

		if _, ok := mapWords[temp]; !ok {
			mapWords[str] = true
		} else {
			result++
		}
	}

	return result
	// res := 0
	// h := make(map[string]bool)
	// for _, w := range words {
	// 	if r := fmt.Sprintf("%c%c", w[1], w[0]); h[r] {
	// 		res++
	// 	} else {
	// 		h[w] = true
	// 	}
	// }
	// return res
}

func reverseString(value string) string {
	runes := []rune(value)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
