package main

import (
	"fmt"
)

func main() {
	var allowed = "abc"
	var words = []string{"a", "b", "c", "ab", "ac", "bc", "abc"}

	fmt.Println(countConsistentStrings(allowed, words))
}

func countConsistentStrings(allowed string, words []string) int {
	mapAllowed := make(map[rune]interface{})

	for _, v := range allowed {
		mapAllowed[v] = struct{}{}
	}

	result := len(words)

	for _, word := range words {
		for _, char := range word {
			if _, ok := mapAllowed[char]; !ok {
				result--
				break
			}
		}
	}

	return result
}
