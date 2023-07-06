package main

func main() {
	test := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(test)
	// fmt.Println(reverseString(test))
}

func reverseString(s []byte) {
	start, end := 0, len(s)-1

	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

// func reverseString(s []byte)  {

//     size := len(s)

//     // reverse string by mirror image
//     for i := 0 ; i < size/2 ; i++{
//         s[i], s[size-1-i] = s[size-1-i], s[i]
//     }

//     return

// }
