package main

import (
	"fmt"
)

func main() {
	var key = "the quick brown fox jumps over the lazy dog"
	var message = "vkbs bs t suepuv"
	// var nums = []int{0, 1, 4, 6, 7, 10}
	// var diff = 3
	fmt.Println(decodeMessage(key, message))
}

func decodeMessage(key string, message string) string {
	mapLetters := make(map[rune]rune, 26)
	result := []rune{}

	var charater rune = 0
	for _, v := range key {
		if _, ok := mapLetters[v]; !ok && v != 32 {
			mapLetters[v] = 'a' + charater
			charater++
		}
	}

	for _, v := range message {
		text := mapLetters[v]
		if v == 32 {
			result = append(result, 32)
		} else {
			result = append(result, text)
		}
	}

	return string(result)
	// table := map[rune]byte{' ': ' '}
	// for _, c := range key {
	// 	if _, ok := table[c]; !ok {
	// 		table[c] = byte(len(table) - 1 + 'a')
	// 	}
	// }

	// var buffer bytes.Buffer
	// for _, c := range message {
	// 	buffer.WriteByte(table[c])
	// }

	// return buffer.String()
}
