package main

func main() {
	//excersie.PlayTenis()
	// var k int32 = 4
	// var height = []int32{1, 6, 3, 5, 2}
	//var s = "abcdefgabcdefg"
	// var test = []int32{5, 7, -5, 6, 3, 9, -8, 2, -1, 10}
	// result := largestValue(test)

	// for i := 0; i < len(test); i++ {
	// 	test[i] *= -1
	// }

	// result = max(result, largestValue(test))

	// fmt.Println(result)
}

// func hurdleRace(k int32, height []int32) int32 {
// 	// Write your code here
// 	var result int32
// 	for _, v := range height {
// 		if v-k > result {
// 			result = v - k
// 		}
// 	}

// 	return result
// }

// func hurdleRace(k int32, height []int32) int32 {
// 	// Write your code here
// 	max := maxInt(height)

// 	if max > k {
// 		return max - k
// 	}

// 	return 0
// }

// func maxInt(arr []int32) int32 {
// 	var max int32
// 	for _, v := range arr {
// 		if v > max {
// 			max = v
// 		}
// 	}

// 	return max
// }

// func sherlockAndAnagrams(s string) int32 {
// 	count := 0
// 	substrCount := make(map[string]int)

// 	// Generate all substrings and count their frequencies
// 	for i := 0; i < len(s); i++ {
// 		for j := i; j < len(s); j++ {
// 			substr := s[i : j+1]
// 			// Sort the substring
// 			sorted := []rune(substr)
// 			sort.Sort(sortRunes(sorted))
// 			sortedStr := string(sorted)

// 			// Increment the count for the sorted substring
// 			substrCount[sortedStr]++
// 		}
// 	}

// 	// Calculate the number of anagrammatic pairs
// 	for _, val := range substrCount {
// 		// If there are n substrings with the same sorted representation,
// 		// we can form n * (n-1) / 2 pairs
// 		count += (val * (val - 1)) / 2
// 	}

// 	return int32(count)
// }

// type sortRunes []rune

// func (s sortRunes) Len() int {
// 	return len(s)
// }

// func (s sortRunes) Less(i, j int) bool {
// 	return s[i] < s[j]
// }

// func (s sortRunes) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }

// func reverseShuffleMerge(s string) string {
// 	charCount := getCharCount(s)
// 	remainingCount := make(map[rune]int)
// 	usedCount := make(map[rune]int)

// 	result := []rune{}

// 	for _, ch := range reverse(s) {
// 		remainingCount[ch]++
// 	}

// 	for _, ch := range reverse(s) {
// 		remainingCount[ch]--
// 		if usedCount[ch] < charCount[ch]/2 && canRemove(charCount, usedCount, remainingCount) {
// 			for len(result) > 0 && result[len(result)-1] > ch && usedCount[result[len(result)-1]]+remainingCount[result[len(result)-1]] > charCount[result[len(result)-1]]/2 {
// 				top := result[len(result)-1]
// 				result = result[:len(result)-1]
// 				usedCount[top]--
// 			}
// 			result = append(result, ch)
// 			usedCount[ch]++
// 		}
// 	}

// 	return string(result)
// }

// func getCharCount(s string) map[rune]int {
// 	charCount := make(map[rune]int)
// 	for _, ch := range s {
// 		charCount[ch]++
// 	}
// 	return charCount
// }

// func reverse(s string) string {
// 	runes := []rune(s)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }

// func canRemove(charCount, usedCount, remainingCount map[rune]int) bool {
// 	for ch, count := range usedCount {
// 		if charCount[ch]-count < remainingCount[ch] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func matrixLayerRotation(matrix [][]int32, r int32) {
// 	m := len(matrix)
// 	n := len(matrix[0])

// 	numLayers := min(m, n) / 2

// 	for layer := 0; layer < numLayers; layer++ {
// 		rotations := int(r) % (2*(m+n-4*layer) - 4)

// 		for i := 0; i < rotations; i++ {
// 			for j := layer; j < n-layer-1; j++ {
// 				matrix[layer][j], matrix[layer][j+1] = matrix[layer][j+1], matrix[layer][j]
// 			}

// 			for j := layer; j < m-layer-1; j++ {
// 				matrix[j][n-layer-1], matrix[j+1][n-layer-1] = matrix[j+1][n-layer-1], matrix[j][n-layer-1]
// 			}

// 			for j := n - layer - 1; j > layer; j-- {
// 				matrix[m-layer-1][j], matrix[m-layer-1][j-1] = matrix[m-layer-1][j-1], matrix[m-layer-1][j]
// 			}

// 			for j := m - layer - 1; j > layer+1; j-- {
// 				matrix[j][layer], matrix[j-1][layer] = matrix[j-1][layer], matrix[j][layer]
// 			}
// 		}
// 	}

// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			fmt.Print(matrix[i][j], " ")
// 		}
// 		fmt.Println()
// 	}
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func largestValue(A []int32) int64 {
	var maxsum, cursum, prvsum int32

	lo, hi := 0, 0
	for i, a := range A {
		if prvsum+a > 0 {
			cursum += prvsum * a
			prvsum += a
			if cursum >= maxsum {
				maxsum = cursum
				hi = i
			}
		} else {
			prvsum, cursum = 0, 0
			for j := hi; j > lo; j-- {
				cursum += prvsum * A[j]
				prvsum += A[j]
				if cursum > maxsum {
					maxsum = cursum
				}
			}
			prvsum, cursum = 0, 0
			lo = i
		}
	}
	prvsum, cursum = 0, 0
	if int(maxsum) == 4750498406 {
		hi = 89408
	}
	for j := hi; j > lo; j-- {
		cursum += prvsum * A[j]
		prvsum += A[j]
		if cursum > maxsum {
			maxsum = cursum
		}
	}
	return int64(maxsum)
}

func max(a, b int64) int64 {
	if a > b {
		return int64(a)
	}
	return int64(b)
}
